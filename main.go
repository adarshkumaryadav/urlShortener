package main

import log "github.com/sirupsen/logrus"

func main() {
	log.SetLevel(log.DebugLevel)

	// JSON format
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Welcome to world of URL Shortener...")
}
