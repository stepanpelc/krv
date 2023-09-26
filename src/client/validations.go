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

//typed-safe client wrap raw CRD client
//comfortable then raw REST-client usage

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"krv/api/crd/v1"
)

type ValidationCrdClientInterface interface {
	List(opts metav1.ListOptions) (*v1.ValidationList, error)
	Get(name string, options metav1.GetOptions) (*v1.Validation, error)
	Update(obj *v1.Validation, options metav1.UpdateOptions) (*v1.Validation, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type validationCrdClient struct {
	restClient rest.Interface
	ns         string
}

func (c *validationCrdClient) List(opts metav1.ListOptions) (*v1.ValidationList, error) {
	result := v1.ValidationList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(v1.CRDPlural).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *validationCrdClient) Get(name string, opts metav1.GetOptions) (*v1.Validation, error) {
	result := v1.Validation{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(v1.CRDPlural).
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *validationCrdClient) Update(obj *v1.Validation, opts metav1.UpdateOptions) (*v1.Validation, error) {
	result := v1.Validation{}
	err := c.restClient.
		Put().
		Name(obj.Name).
		Namespace(c.ns).
		Resource(v1.CRDPlural).
		Body(obj).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (c *validationCrdClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource(v1.CRDPlural).
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.TODO())
}
