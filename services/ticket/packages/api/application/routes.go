package application

import (
	"github.com/dindasigma/go-microservices-ticket/packages/api/controllers"
)

func initializeRoutes() {
	// Ticket Route
	router.HandleFunc("/ticket", controllers.TicketController.Index).Methods("GET")
}
