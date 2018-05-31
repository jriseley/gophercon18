package main

import (
	"log"
	"net/http"
	"github.com/jriseley/gophercon18/pkg/routing"
)


func main() {
	log.Printf("Service is starting...")
	r := routing.BaseRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}



// To test the service: 
// go run ./cmd/gophercon/gophercon.go
// curl -i http://127.0.0.1:8000/home



