package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	// JSON format
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Welcome to world of URL Shortener...")
	// let's start the server
	server := http.Server{
		Addr: ":8080",
		// just in case app is taking more time to process the request
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}
