/*
    krv - kubernetes resource validator
    Copyright (C) 2022 SIZEK s.r.o

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package watcher

//in time period iterate over Validation resources and by .spec definition try to find corresponding kubernetes resource

import (
	"fmt"
	"krv/shared"
	"os"
	"strconv"
	"time"
)

//periodically get all Validation resources and check if conditions are met
//actualddize resource state

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/json"
	v1 "krv/api/crd/v1"
	"krv/client"
	"krv/service/crd"
	"regexp"
	"strings"
)

var timePeriodInMinutes = 5 //default
var apiVersionsMap map[string][]string

func init() {
	if os.Getenv("TIME_PERIOD") != "" {
		timePeriodInMinutes, _ = strconv.Atoi(os.Getenv("TIME_PERIOD"))
	}
}

func Start() {
	ticker := time.NewTicker(time.Duration(timePeriodInMinutes) * time.Minute)
	log.Debug().Msgf("Time interval set top %s minutes", time.Duration(timePeriodInMinutes)*time.Minute)
	quit := make(chan struct{})
	shared.HealthStatus = shared.HealthOK
	check()
	for {
		select {
		case <-ticker.C:
			check()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func check() {
	log.Info().Msgf("Checking validations")
	validations, err := crd.GetValidationList()
	//update available api versions of k8s resources in cluster
	apiVersionsMap = client.GetApiGroupsVersions()
	validationNamespaceMap := make(map[string][]*v1.Validation)
	if err != nil {
		log.Error().Msgf("Cannot get ValidationList: %v", err)
	} else {
		log.Debug().Msgf("Found %v Validation(s)", len(validations.Items))

		//the goal is to minimize api-server request hits,
		//so we get every resource in batch once and store it in map
		//later than, already stored resource are used

		//group CRD Validation resources by its spec.namespace
		for i, validation := range validations.Items {
			validationNamespaceMap[validation.Spec.Namespace] = append(validationNamespaceMap[validation.Spec.Namespace], &validations.Items[i])
		}
		log.Trace().Msgf("Validation definitions mapped by namespace: %s", validationNamespaceMap)

		//iterate over all Validations for given .spec.namespace
		for ns, validationList := range validationNamespaceMap {

			log.Debug().Msgf("For %s namespace there are %v Validation(s) defined", ns, len(validationList))
			log.Trace().Msgf("%s: %s", ns, validationList)

			//for every new k8s resource type all resources are requested and store into map
			//the map key is resourceType and the value is ListOfResources
			resourceTypeMap := make(map[string]*unstructured.UnstructuredList)
			// go over all validations in actual .spec.namespace
			for _, validation := range validationList {

				log.Debug().Msgf("Process %s Validation", validation.Name)

				//store actual k8s resource to map (if it is not already there)
				addAllByResourceTypeToMap(validation, resourceTypeMap)

				//prepare new Validation status values
				oldState := validation.Status.State
				timeNow := time.Now().Format("2006-01-02 15:04:05")
				log.Debug().Msgf("Set initial MISSING status for %s before appropriate resources are found and validated", validation.Name)
				validation.Status.State = "MISSING"
				validation.Status.LastCheck = timeNow

				log.Debug().Msgf("Look for %s resource in namespace %s with name pattern %s", validation.Spec.Resource, validation.Spec.Namespace, validation.Spec.Name)

				//load all validations for actual k8s resource type from map
				unstructedResourcelist := resourceTypeMap[resourcePluralName(validation.Spec.Resource)]
				if unstructedResourcelist != nil {
					//go throw resources and find the correct one by name and namespace
					for _, unstructedResource := range unstructedResourcelist.Items {
						log.Trace().Msgf("Check namespace %s match %s and name %s match %s", unstructedResource.GetNamespace(), validation.Spec.Namespace, unstructedResource.GetName(), validation.Spec.Name)
						validatedResourceNameMatch, _ := regexp.MatchString(validation.Spec.Name, unstructedResource.GetName())
						if unstructedResource.GetNamespace() == validation.Spec.Namespace && validatedResourceNameMatch {
							log.Debug().Msgf("Found match resource: %s", unstructedResource.GetName())
							checkValidationRules(validation, &unstructedResource)

							//if validation of the current resource end up with NOK status it means whole Validation resource end up with NOK status
							//and there is no need to found and validate another match resource for current Validation resource
							//we skip out from the for loop
							if validation.Status.State == "NOK" {
								break
							}
						}
					}
				}

				//if status changed
				if validation.Status.LastChanged == "" || oldState != validation.Status.State {
					validation.Status.LastChanged = timeNow
				}
				//update k8s Validation resource
				_, err = crd.UpdateValidation(validation, oldState != validation.Status.State)
				if err != nil {
					log.Error().Msgf("%v", err.Error())
				}
			}
		}
	}
}

// ensure given resource is in plural format
func resourcePluralName(name string) string {
	name = strings.ToLower(name)
	if match, _ := regexp.MatchString("^.+s$", name); match {
		return name
	}
	return name + "s"
}

// retrieve all resource of type specified in validation.spec.resource and store in into map
func addAllByResourceTypeToMap(validation *v1.Validation, resourceTypeMap map[string]*unstructured.UnstructuredList) {
	var resourcePluaralName = resourcePluralName(validation.Spec.Resource)
	if resourceTypeMap[resourcePluaralName] == nil {
		log.Debug().Msgf("Get all %s resources in %s namespace", resourcePluaralName, validation.Spec.Namespace)
		var errs = ""
		//check we have appropriate api version for current k8s resource
		if len(apiVersionsMap[resourcePluaralName]) > 0 {
			//check all available api versions for current k8s resource. If success, break the iteration
			for _, apiVersion := range apiVersionsMap[resourcePluaralName] {
				apiGroup := ""
				if strings.Count(apiVersion, "/") > 0 {
					split := strings.Split(apiVersion, "/")
					apiGroup = split[0]
					apiVersion = split[1]
				}

				nodeResource := schema.GroupVersionResource{Resource: resourcePluaralName, Version: apiVersion, Group: apiGroup}
				unstructlist, err := client.DynamicClientSet.Resource(nodeResource).Namespace(validation.Spec.Namespace).List(context.TODO(), metav1.ListOptions{})
				if err != nil {
					resourceTypeMap[resourcePluaralName] = nil
					errs = fmt.Sprintf("DynamicClient cannot get list of %s in %s namespace: %v .", nodeResource, validation.Spec.Namespace, err.Error())
				} else {
					resourceTypeMap[resourcePluaralName] = unstructlist
					//null the possible previous errors
					errs = ""
					break
				}
			}
			//if client was not able to get the resource, log error messages
			if errs != "" {
				log.Warn().Msgf(errs)
			}
		} else {
			log.Warn().Msgf("Cannot investigate api version for resource %s. Validation of %s validation definition skipped", resourcePluaralName, validation.Name)
		}
	}
}

// for given Validation resource it iterates over adequate k8s resources  and looks for the one that match the rules
func checkValidationRules(validation *v1.Validation, unstructedResource *unstructured.Unstructured) {
	log.Debug().Msgf("Check validation rules for %s %s/%s", validation.Spec.Namespace, unstructedResource.GetNamespace(), unstructedResource.GetName())
	//we found the resource for now status is OK
	validation.Status.State = "OK"
	defer log.Debug().Msgf("Validation status for %s set to %s", validation.Name, validation.Status.State)
	// Convert our unstructured resource object to raw JSON
	var rawJsonInterface interface{}
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(unstructedResource.Object, &rawJsonInterface)
	if err != nil {
		log.Error().Msgf("Cannot convert from unstructured object to interface: %v", err.Error())
		//if we cannot parse Unstructured object lets handle it as NOK status
		validation.Status.State = "NOK"
	}
	rawJson, err := json.Marshal(rawJsonInterface)
	if err != nil {
		//if we cannot parse marshall json object lets handle it as NOK status
		validation.Status.State = "NOK"
	}
	jsonString := string(rawJson[:])
	log.Trace().Msgf("Actual validated resource: %s", jsonString)
	//now lets check the rules by jq and if it is not fit set status to NOK and quit
	for _, rule := range validation.Spec.Validation {
		value := gjson.Get(jsonString, rule.JsonPath).Str
		if match, _ := regexp.MatchString(rule.Value, value); !match {
			//if condition not met let set NOK status
			validation.Status.State = "NOK"
			break
		}
	}
}
