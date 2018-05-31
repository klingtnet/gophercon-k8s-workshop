package main

import (
	"log"
	"os"
	"sync"

	"github.com/klingtnet/gophercon/pkg/routing"
	"github.com/klingtnet/gophercon/pkg/webserver"
	"github.com/klingtnet/gophercon/version"
)

// go run ./cmd/gophercon/gophercon.go
// curl -i http://127.0.0.1:8000/home
func main() {
	log.Printf(
		"Service is starting, version is %s, commit is %s, time is %s...",
		version.Release, version.Commit, version.BuildTime,
	)

	// you can also use github.com/kelseyhightower/envconfig
	// to keep your config more structured
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Service port wasn't set")
	}

	r := routing.BaseRouter()
	ws := webserver.New("", port, r)

	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		ws.Start()
		w.Done()
		log.Printf("Gracefully shutting down web server...")
	}()

	internalPort := os.Getenv("INTERNAL_PORT")
	if len(internalPort) == 0 {
		log.Fatal("Internal port wasn't set")
	}
	diagnosticsRouter := routing.DiagnosticsRouter()
	diagnosticsServer := webserver.New(
		"", internalPort, diagnosticsRouter,
	)

	w.Add(1)
	diagnosticsServer.Start()
	w.Done()
	log.Printf("Gracefully shutting down diagnostics server...")
}
