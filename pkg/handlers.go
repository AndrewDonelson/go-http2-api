package api

import (
	"fmt"
	"net/http"
)

// HANDLER FUNCTIONS (PUBLIC) - Add your handlers here

// HeartbeatHandler checks the health of the service. For simplicity, we'll just send an "OK" response.
func HeartbeatHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Service is alive!")
}
