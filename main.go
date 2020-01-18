package main

import "net/http"

func main() {
	router := ServeRoutes()
	http.ListenAndServe(":444", router)
}
