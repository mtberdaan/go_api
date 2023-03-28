package main

import (
	"log"
	"net/http"

	"github.com/mtberdaan/go_api/go_api"
)

func main() {
	router := go_api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
