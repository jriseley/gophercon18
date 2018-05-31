package main

import (
	"log"
	"github.com/jriseley/gophercon18/pkg/routing"
	"github.com/jriseley/gophercon18/pkg/webserver"
	"os"
)


func main() {
	log.Printf("Service is starting...")
	port := os.Getenv("SERVICE_PORT")

	if len(port) == 0 {
		log.Fatal("Service port wasn't set\n")
	}

	r := routing.BaseRouter()
	ws := webserver.New("", port, r)

	log.Fatal(ws.Start())
}



// To test the service: 
// go run ./cmd/gophercon/gophercon.go
// curl -i http://127.0.0.1:8000/home



