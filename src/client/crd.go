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

package client

//create client instance for our Validations CRD

import (
	"github.com/rs/zerolog/log"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	crdapiv1 "krv/api/crd/v1"
)

var CrdClientset *ValidationCrdV1Client = nil

type ValidationV1CrdInterface interface {
	Validations(namespace string) ValidationCrdClientInterface
}

type ValidationCrdV1Client struct {
	restClient rest.Interface
}

func init() {
	CrdClientset = newCRDClient()
}

// create CRD client instance
func newCRDClient() *ValidationCrdV1Client {

	if config, err := rest.InClusterConfig(); err != nil {
		log.Error().Msgf("Cannot get InClusterConfig %v", err.Error())
		panic(err)
	} else {
		config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: crdapiv1.CRDGroup, Version: crdapiv1.CRDVersion}
		config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: crdapiv1.CRDGroup, Version: crdapiv1.CRDVersion}
		config.APIPath = "/apis"
		config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
		config.UserAgent = rest.DefaultKubernetesUserAgent()
		client, err := rest.RESTClientFor(config)
		if err != nil {
			log.Error().Msgf("Cannot initialize CRD client, %v", err.Error())
			panic(err)
		}
		return &ValidationCrdV1Client{restClient: client}
	}
}

func (c *ValidationCrdV1Client) Validations(namespace string) ValidationCrdClientInterface {
	return &validationCrdClient{
		restClient: c.restClient,
		ns:         namespace,
	}
}
