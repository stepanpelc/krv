package watcher

//watch our Validation CRDs and save the latest state in cache

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	v1 "krv/api/crd/v1"
	crdservice "krv/service/crd"
	"time"
)

func WatchResources() cache.Store {
	validatiionStore, projectController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return crdservice.GetValidationList()
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return crdservice.WatchValidations()

			},
		},
		&v1.Validation{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{},
	)

	go projectController.Run(wait.NeverStop)
	return validatiionStore
}
