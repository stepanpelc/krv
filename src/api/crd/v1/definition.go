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
