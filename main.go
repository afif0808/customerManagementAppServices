package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
	router := ServeRoutes()
	http.ListenAndServe(":444", cors.AllowAll().Handler(router))
}
