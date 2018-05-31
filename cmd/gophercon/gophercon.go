package main

import (
	"log"
	"net/http"

	"github.com/Subsack/gophercon/pkg/routing"
)

func main() {
	log.Printf("Service is starting")

	router := routing.BaseRouter()
	//http.HandleFunc("/", handler())
	log.Fatal(http.ListenAndServe(":8000", router))
}
