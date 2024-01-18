package main

import (
	"net/http"
)

func main() {
	// function name should be start with capital letter
	setupJsonApi()
	// add NewRoute() to return new route instance and add healthcheck() to add liveness and readness
	http.ListenAndServe(":80", nil)
}
