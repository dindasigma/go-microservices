package controllers

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/dindasigma/go-microservices-messaging/datasources"
	"github.com/dindasigma/go-microservices-messaging/models/users"
	"github.com/nsqio/go-nsq"
)

var (
	EmailController emailControllerInterface = &emailController{}
)

type emailControllerInterface interface {
	SendWelcome(*nsq.Message) error
}

type emailController struct{}

func (c emailController) SendWelcome(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		// In this case, a message with an empty body is simply ignored/discarded.
		return nil
	}

	// do whatever actual message processing is desired
	//err := processMessage(m.Body)
	uid, err := strconv.ParseUint(string(m.Body), 10, 32)
	if err != nil {
		log.Print(err)
		return errors.New("failed to convert type")
	}

	user := users.User{}
	userGotten, err := user.FindByID(datasources.DB, uint32(uid))
	if err != nil {
		// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
		log.Print(err)
		return errors.New("userid not found")
	}

	emailBody := fmt.Sprintf("halo %s %s", userGotten.FirstName, userGotten.LastName)
	log.Printf("WE ARE SENDING EMAIL WELCOME TO CUSTOMER WITH BODY %s", emailBody)

	return nil
}
