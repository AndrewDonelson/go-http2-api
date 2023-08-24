package api

import (
	"testing"
)

func TestServerInitialize(t *testing.T) {
	Server.Initialize(":8080", "../certs")
	if Server.GetPort() != ":8080" {
		t.Errorf("expected port :8080, got %s", Server.GetPort())
	}
	if Server.GetCertFolder() != "../certs" {
		t.Errorf("expected cert folder ../certs, got %s", Server.GetCertFolder())
	}
}
