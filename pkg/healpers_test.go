package api

import (
	"testing"
)

func TestGetCertPaths(t *testing.T) {
	// Change this for your repo
	basePath := "/home/andrew/go/src/github.com/AndrewDonelson/go-http2-api"

	certFolder := "../certs/"
	certPath, keyPath := getCertPaths(certFolder)

	expectedCertPath := basePath + "/certs/cert.pem"
	expectedKeyPath := basePath + "/certs/key.pem"

	if certPath != expectedCertPath {
		t.Errorf("expected cert path %s, got %s", expectedCertPath, certPath)
	}

	if keyPath != expectedKeyPath {
		t.Errorf("expected key path %s, got %s", expectedKeyPath, keyPath)
	}
}

// You can add more tests for other helper functions here.
