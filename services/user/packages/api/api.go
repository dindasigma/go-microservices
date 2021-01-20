package api

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dindasigma/go-microservices-user/packages/api/controllers"
	"github.com/dindasigma/go-microservices-user/packages/api/datasources"
	"github.com/dindasigma/go-microservices-user/packages/api/servers"
	"github.com/joho/godotenv"
)

const (
	grpcPort = "50051"
	restPort = "8080"
)

// app is a convenience wrapper for all things needed to start
// and shutdown the Order microservice
type app struct {
	restServer servers.RestServer
	grpcServer servers.GrpcServer
	/* Listens for an application termination signal
	   Ex. (Ctrl X, Docker container shutdown, etc) */
	shutdownCh chan os.Signal
}

// get env and connect db
func (a app) prepare() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	datasources.InitializePostgres()

}

// start starts the REST and gRPC Servers in the background
func (a app) start() {
	a.restServer.Start() // non blocking now
	a.grpcServer.Start() // also non blocking :-)
}

// stop shuts down the servers
func (a app) shutdown() error {
	a.grpcServer.Stop()
	return a.restServer.Stop()
}

// newApp creates a new app with REST & gRPC servers
// this func performs all app related initialization
func newApp() (app, error) {
	userService := controllers.GrpcServer{}

	gs, err := servers.NewGrpcServer(userService, grpcPort)
	if err != nil {
		return app{}, err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	return app{
		restServer: servers.NewRestServer(restPort),
		grpcServer: gs,
		shutdownCh: quit,
	}, nil
}

// run starts the app, handling any REST or gRPC server error
// and any shutdown signal
func Run() error {
	app, err := newApp()
	if err != nil {
		return err
	}

	app.prepare()
	app.start()
	defer app.shutdown()

	select {
	case restErr := <-app.restServer.Error():
		return restErr
	case grpcErr := <-app.grpcServer.Error():
		return grpcErr
	case <-app.shutdownCh:
		return nil
	}
}
