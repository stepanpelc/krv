package http

//return latest state of all Validation resources.
//for that purpose cache is used

import (
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/tools/cache"
	"krv/watcher"
	"net/http"
)

var cacheStore cache.Store

func init() {
	cacheStore = watcher.WatchResources()
}

func GetAllValidations(w http.ResponseWriter, _ *http.Request) {
	b, err := json.Marshal(cacheStore.List())
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}
	w.Write(b)
	w.WriteHeader(200)
}
