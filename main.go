package main

import (
	"log"
	"net/http"

	"github.com/rrgaya/mock/api"
)

func main() {

	http.HandleFunc("/", api.GetRandomUser)

	log.Println("Lintening in 8000...")
	panic(http.ListenAndServe("localhost:8000", nil))
}
