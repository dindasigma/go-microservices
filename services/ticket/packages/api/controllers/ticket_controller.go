package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	TicketController ticketControllerInterface = &ticketController{}
)

type ticketControllerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type ticketController struct{}

func (c *ticketController) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("Welcome to The Ticket API")
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
