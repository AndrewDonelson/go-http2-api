package api

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
)

// HELPER FUNCTIONS (PRIVATE)

// getCertPaths returns the paths to the TLS certificates.
func getCertPaths(certFolder string) (string, string) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	certPath := filepath.Join(basepath, certFolder+"cert.pem")
	keyPath := filepath.Join(basepath, certFolder+"key.pem")
	return certPath, keyPath
}

func waitForShutdown() {
	// Create a channel to listen for OS signals
	signals := make(chan os.Signal, 1)

	// Notify the signals channel for SIGINT (CTRL+C) and SIGTERM (termination request sent to the program)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive the signal
	<-signals

	log.Println("Shutdown signal received")
}

// cleanup closes the server when the application exits.
func cleanup() {
	Server.Stop()
	log.Println("Server stopped")
}

// StringArrayContains helper function to return true or false is a given string exists in the provided string array
func StringArrayContains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
