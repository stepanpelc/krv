package v1

// Validation CRD openapi description

import (
	apiextensionv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
)

const CRDPlural string = "validations"
const CRDSingular string = "validation"
const FullCRDName = CRDPlural + "." + CRDGroup

var ValidationCRDDefinition = &apiextensionv1.CustomResourceDefinition{
	ObjectMeta: meta_v1.ObjectMeta{Name: FullCRDName},
	Spec: apiextensionv1.CustomResourceDefinitionSpec{
		Group: CRDGroup,
		Versions: []apiextensionv1.CustomResourceDefinitionVersion{
			{
				Name: CRDVersion,
				Schema: &apiextensionv1.CustomResourceValidation{
					OpenAPIV3Schema: &apiextensionv1.JSONSchemaProps{
						Type: "object",
						Properties: map[string]apiextensionv1.JSONSchemaProps{
							"spec": {
								Type: "object",
								Properties: map[string]apiextensionv1.JSONSchemaProps{
									"name": {
										Type: "string",
									},
									"namespace": {
										Type: "string",
									},
									"resource": {
										Type: "string",
									},
									"validation": {
										Type: "array",
										Items: &apiextensionv1.JSONSchemaPropsOrArray{
											Schema: &apiextensionv1.JSONSchemaProps{
												Type: "object",
												Properties: map[string]apiextensionv1.JSONSchemaProps{
													"jsonPath": {
														Type: "string",
													},
													"value": {
														Type: "string",
													},
												},
											},
										},
									},
								},
								Required: []string{"name", "namespace", "resource"},
							},
							"status": {
								Type: "object",
								Properties: map[string]apiextensionv1.JSONSchemaProps{
									"lastCheck": {
										Type: "string",
									},
									"lastChange": {
										Type: "string",
									},
									"state": {
										Type: "string",
									},
								},
								Default: nil,
							},
						},
						Required: []string{"spec"},
					},
				},
				Storage: true,
				Served:  true,
				AdditionalPrinterColumns: []apiextensionv1.CustomResourceColumnDefinition{
					{
						Name:        "RESOURCE-NAMESPACE",
						Type:        "string",
						Description: "Namespace of validated resource",
						JSONPath:    ".spec.namespace",
					},
					{
						Name:        "RESOURCE",
						Type:        "string",
						Description: "Type of validated resource",
						JSONPath:    ".spec.resource",
					},
					{
						Name:        "STATE",
						Type:        "string",
						Description: "Actual state of validated resource",
						JSONPath:    ".status.state",
					},
					{
						Name:     "AGE",
						Type:     "date",
						JSONPath: ".metadata.creationTimestamp",
					},
				},
			},
		},
		Scope: apiextensionv1.NamespaceScoped,
		Names: apiextensionv1.CustomResourceDefinitionNames{
			Singular:   CRDSingular,
			ShortNames: []string{"val", "vals"},
			Plural:     CRDPlural,
			Kind:       reflect.TypeOf(Validation{}).Name(),
			Categories: []string{"all"},
		},
	},
}
