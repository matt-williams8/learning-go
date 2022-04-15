package main

import (
	"dependency-injection/api"
	"log"
	"net/http"
)

func GreeterHandler(responseWriter http.ResponseWriter, request *http.Request) {
	api.Greet(responseWriter, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(GreeterHandler)))
}
