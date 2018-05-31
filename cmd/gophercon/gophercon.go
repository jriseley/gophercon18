package main

import (
	"log"
	"net/http"
	"github.com/jriseley/gophercon18/pkg/routing"
	"os"
)


func main() {
	log.Printf("Service is starting...")
	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Service port wasn't set\n")
	}
	r := routing.BaseRouter()
	log.Fatal(http.ListenAndServe(":"+port, r))
}



// To test the service: 
// go run ./cmd/gophercon/gophercon.go
// curl -i http://127.0.0.1:8000/home



