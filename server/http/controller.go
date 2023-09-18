package http

//simple http controller to handle incoming requests

import (
	"github.com/rs/zerolog/log"
	"krv/shared"
	"net/http"
)

func Start() {
	http.HandleFunc("/health", Health)
	http.HandleFunc("/validations", GetAllValidations)
	go http.ListenAndServe(":"+shared.PORT, nil)
	log.Info().Msgf("Start listening on port %v", shared.PORT)
}
