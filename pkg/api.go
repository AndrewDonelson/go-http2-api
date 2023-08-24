package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// APIRoute is a struct to represent an API route.
type APIRoute struct {
	Version  uint8                                    // 0 = v1, 1 = v2, etc.
	SubRoute string                                   // e.g. "users", "posts","products", etc.
	Name     string                                   // e.g. "get", "create", "update", "delete", etc.
	Method   string                                   // e.g. "GET", "POST", "PUT", "DELETE", etc.
	Handler  func(http.ResponseWriter, *http.Request) // e.g. GetUsersHandler, CreateUsersHandler, etc.
}

// APIServer is a struct to represent the API server.
type APIServer struct {
	port              string       // e.g. ":8080"
	certFolder        string       // e.g. "./certs"
	server            *http.Server // the server object
	router            *mux.Router  // the router object
	authorizedRemotes []string     // the list of authorized remotes (IP Addresses)
}

var (
	// Server is the global value for the API server
	Server *APIServer
)

// init initializes and empty API server whn the package is loaded
func init() {
	Server = &APIServer{}
}

// AddRoute adds a route to the API server.
// Example:
//
//	Server.AddRoute(&APIRoute{
//		version:  1,
//		subRoute: "users",
//		name:     "get",
//		method:   "GET",
//		handler:  GetUsersHandler,
//	})
func (s *APIServer) AddRoute(route *APIRoute) {
	s.router.HandleFunc(fmt.Sprintf("/v%d/%s/%s", route.Version, route.SubRoute, route.Name), route.Handler).Methods(strings.ToUpper(route.Method))
}

// AddAuthorizedRemote adds an authorized remote to the API server.
func (s *APIServer) AddAuthorizedRemote(remote string) {
	s.authorizedRemotes = append(s.authorizedRemotes, remote)
}

// startServer starts the HTTP/2 server in a goroutine.
func (s *APIServer) Initialize(port string, certFolder string) {
	s.port = port
	s.certFolder = certFolder

	log.Println("Configuring router")
	s.router = mux.NewRouter()

	// Heartbeat Endpoint (/):
	// This endpoint is designed to check the health of the service.
	// Details:
	// - Purpose: To check the health of the service.
	// - Functionality: The endpoint returns a simple "OK" response.
	// - Output: A string indicating that the service is alive.
	// - Use Case: This endpoint can be used to check the health of the service.
	//
	// TODO: Add Version Information, Database Connection Status, Persistent Storage Status, etc.
	s.router.HandleFunc("/", HeartbeatHandler).Methods("GET")

	s.server = &http.Server{
		Addr:    port,
		Handler: s.router,
	}

}

// Start starts the HTTP/2 server in a goroutine. It then calls waitForShutdown() to block until the server is stopped.
func (s *APIServer) Start() {

	defer cleanup()

	// Start the server in a goroutine so that it doesn't block.
	go func() {
		// Usage in your server setup:
		log.Println("Loading TLS certificates")
		certPath, keyPath := getCertPaths(s.certFolder)

		log.Printf("http/2 API Server listening on %s\n", s.server.Addr)
		if err := s.server.ListenAndServeTLS(certPath, keyPath); err != nil {
			log.Fatal(err)
		}
	}()

	waitForShutdown()
}

// Stop stops the HTTP/2 server.
func (s *APIServer) Stop() {
	//if we have a server, stop it
	if s.server != nil {
		log.Default().Println("Stopping server")
		s.server.Close()
	}
}
