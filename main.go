package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	database := NewDBConnection()
	fmt.Println("db connected")
	fmt.Println("api up at 8080")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
