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

	go func() {
		log.Fatal(ws.Start())
	}()
	
	internalPort := os.Getenv("INTERNAL_PORT")

	
	if len(internalPort) == 0 {
		log.Fatal("Internal port wasn't set\n")
	}

	diagnosticsRouter := routing.DiagonsticsRouter()
	diagnosticsServer := webserver.New(
	"", internalPort, diagnosticsRouter,
	)
	log.Fatal(diagnosticsServer.Start())
}



// To test the service: 
// go run ./cmd/gophercon/gophercon.go
// curl -i http://127.0.0.1:8000/home



