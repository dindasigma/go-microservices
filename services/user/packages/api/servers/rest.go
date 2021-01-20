package servers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RestServer implements a REST server for the Order Service
type RestServer struct {
	server *http.Server
	errCh  chan error
}

var (
	router = mux.NewRouter()
)

// NewRestServer is a convenience func to create a RestServer
func NewRestServer(port string) RestServer {

	rs := RestServer{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
		errCh: make(chan error),
	}

	// register routes
	initializeRoutes()

	return rs
}

// Start starts the server
func (r RestServer) Start() {
	go func() {
		r.errCh <- r.server.ListenAndServe()
	}()
}

// Stop stops the server
func (r RestServer) Stop() error {
	return r.server.Close()
}

// Error returns the server's error channel
func (r RestServer) Error() chan error {
	return r.errCh
}
