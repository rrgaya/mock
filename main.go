package main

import (
	"log"
	"net/http"

	"github.com/rrgaya/mock/api"
)

func main() {

	service := api.Service{
		HTTPClient: http.DefaultClient,
	}

	http.HandleFunc("/", service.GetRandomUser)

	log.Println("Listening localhost:8000")
	panic(http.ListenAndServe("localhost:8000", nil))
}
