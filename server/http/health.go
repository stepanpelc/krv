package http

//hold and return actual health status

import (
	"krv/shared"
	"net/http"
)

func Health(w http.ResponseWriter, _ *http.Request) {
	if shared.HealthStatus {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(500)
		w.Write([]byte("NOK"))
	}
}
