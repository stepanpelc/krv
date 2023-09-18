package client

//create standard kubernetes clients instances
//we need classic kubernetes Clientset, apiextension and dynamic client

import (
	"github.com/rs/zerolog/log"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"krv/api/crd/v1"
)

var Clientset *kubernetes.Clientset
var ApiExtensionClientset apiextension.Interface
var DynamicClientSet dynamic.Interface
var discoveryClient discovery.DiscoveryInterface

func init() {
	v1.AddToScheme(scheme.Scheme) //register custom types
	if config, err := rest.InClusterConfig(); err != nil {
		log.Error().Msgf("Cannot get InClusterConfig %v", err.Error())
		panic(err)
	} else {

		if DynamicClientSet, err = dynamic.NewForConfig(config); err != nil {
			log.Error().Msgf("Cannot initialize Kubernetes Dynamic Clientset, %v", err.Error())
			panic(err)
		}

		if Clientset, err = kubernetes.NewForConfig(config); err != nil {
			log.Error().Msgf("Cannot initialize Kubernetes Clientset, %v", err.Error())
			panic(err)
		}

		ApiExtensionClientset, err = apiextension.NewForConfig(config)
		if err != nil {
			log.Error().Msgf("Cannot to create ApiExtension Clientset: %v", err.Error())
			panic(err)
		}

		if discoveryClient, err = discovery.NewDiscoveryClientForConfig(config); err != nil {
			log.Error().Msgf("Cannot initialize Kubernetes Dynamic Clientset, %v", err.Error())
			panic(err)
		}
	}
}

//GetApiGroupsVersions return map of k8s resource name and its api version
func GetApiGroupsVersions() map[string][]string {

	var apiVersionsMap = make(map[string][]string)
	_, apiResourceListArray, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		log.Error().Msgf("Unable to get server groups and resources: %v", err.Error())
		return apiVersionsMap
	}
	for _, apiResourceList := range apiResourceListArray {
		for _, apiResource := range apiResourceList.APIResources {
			apiVersionsMap[apiResource.Name] = append(apiVersionsMap[apiResource.Name], apiResourceList.GroupVersion)
		}
	}
	log.Debug().Msgf("Retrieved resources api versions: %s", apiVersionsMap)
	return apiVersionsMap
}
