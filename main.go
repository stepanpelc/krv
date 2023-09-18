package main

//entrypoint of the whole application

import (
	"krv/server/http"
	"krv/watcher"
)

func main() {
	http.Start()
	watcher.Start()
}
