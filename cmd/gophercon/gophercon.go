package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/jriseley/gophercon18/pkg/routing"
	"github.com/jriseley/gophercon18/pkg/webserver"
	"github.com/jriseley/gophercon18/version"

)


func main() {
	log.Printf("Service is starting, version is %s...", version.Release)

	shutdown := make(chan error, 2)

	port := os.Getenv("PORT")

	if len(port) == 0 {
		log.Fatal("Service port wasn't set\n")
	}

	r := routing.BaseRouter()
	ws := webserver.New("", port, r)

	go func() {
		err := ws.Start()
		shutdown <- err
	}()
	
	internalPort := os.Getenv("INTERNAL_PORT")

	if len(internalPort) == 0 {
		log.Fatal("Internal port wasn't set\n")
	}

	diagnosticsRouter := routing.DiagonsticsRouter()
	diagnosticsServer := webserver.New(
	"", internalPort, diagnosticsRouter,
	)

	go func() {
		err := diagnosticsServer.Start()
		shutdown <- err
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case killSignal:=<-interrupt:
		log.Printf("Got %s. Stopping...\n", killSignal)
	case err:=<-shutdown:
		log.Printf("Got an error '%s'. Stopping\n", err)
	}

	err := ws.Stop()
	if err != nil {
		log.Print(err)
	}

	err = diagnosticsServer.Stop()

	if err != nil {
		log.Print(err)
	}
	// stop extra tasks
	//os.Exit(...)

}



// To test the service: 
// go run ./cmd/gophercon/gophercon.go
// curl -i http://127.0.0.1:8000/home



