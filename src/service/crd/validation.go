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

package crd

//service layer over CRD resources operations.
//so no need to use the client directly

import (
	"context"
	"krv/client"
	"time"

	"github.com/rs/zerolog/log"
	"k8s.io/client-go/tools/record"

	core_v1 "k8s.io/api/core/v1"

	apierrors "k8s.io/apimachinery/pkg/api/errors"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"

	tyoed_core_v1 "k8s.io/client-go/kubernetes/typed/core/v1"

	v1 "krv/api/crd/v1"
	"krv/shared"
)

var recorder record.EventRecorder

func init() {
	createValidationCRD()
	recorder = createEventRecorder()
}

func createValidationCRD() {
	_, err := client.ApiExtensionClientset.
		ApiextensionsV1().
		CustomResourceDefinitions().
		Create(context.TODO(), v1.ValidationCRDDefinition, meta_v1.CreateOptions{})
	if err != nil && apierrors.IsAlreadyExists(err) {
		log.Info().Msg("CRD already exist. No need to register it again")
		return
	} else if err != nil {
		panic("CRD registration failed")
	}
	//wait till new CRD is available
	for {
		_, err := client.ApiExtensionClientset.ApiextensionsV1().CustomResourceDefinitions().Get(context.TODO(), v1.FullCRDName, meta_v1.GetOptions{})
		if err == nil {
			break
		}
		log.Trace().Msg("CRD not available yet")
		time.Sleep(100 * time.Millisecond)
	}
	log.Info().Msg("CRD successfully registered")
}

func GetValidationList() (*v1.ValidationList, error) {
	return client.CrdClientset.Validations(shared.KrvNs).List(meta_v1.ListOptions{})
}

func UpdateValidation(resource *v1.Validation, sentEvent bool) (*v1.Validation, error) {
	newVal, err := client.CrdClientset.Validations(shared.KrvNs).Update(resource, meta_v1.UpdateOptions{})
	if err == nil && sentEvent {
		recorder.Event(newVal, "Normal", "StateChanged", newVal.Status.State)
	}
	return newVal, err
}

func WatchValidations() (watch.Interface, error) {
	return client.CrdClientset.Validations(shared.KrvNs).Watch(meta_v1.ListOptions{})
}

func createEventRecorder() record.EventRecorder {
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(log.Trace().Msgf)
	eventBroadcaster.StartRecordingToSink(
		&tyoed_core_v1.EventSinkImpl{
			Interface: client.Clientset.CoreV1().Events(shared.KrvNs)})
	return eventBroadcaster.NewRecorder(scheme.Scheme, core_v1.EventSource{Component: "krv"})
}
