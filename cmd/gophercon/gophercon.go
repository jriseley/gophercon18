package main

import (
	"log"
	"os"
	"github.com/jriseley/gophercon18/pkg/routing"
	"github.com/jriseley/gophercon18/pkg/webserver"
	"github.com/jriseley/gophercon18/version"
)


func main() {
	log.Printf("Service is starting, version is %s...", version.Release)
	port := os.Getenv("PORT")

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



