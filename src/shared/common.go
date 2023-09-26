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
