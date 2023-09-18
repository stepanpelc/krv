package shared

//common values and constants shared across all packages

import (
	"github.com/rs/zerolog"
	"os"
)

const HealthOK = true
const HealthNOK = false

var HealthStatus = HealthNOK

const PORT = "8080"

var KrvNs = os.Getenv("NAMESPACE")
var logLevel, _ = zerolog.ParseLevel("info")

func init() {
	if os.Getenv("LOG_LEVEL") != "" {
		logLevel, _ = zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
	}
	zerolog.SetGlobalLevel(logLevel)
}
