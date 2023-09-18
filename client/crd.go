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
