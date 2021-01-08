package controllers

import (
	"log"

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
	userid := string(m.Body)
	log.Printf("we are sending email welcome to customer with id %s", userid)
	return nil

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	//return errors.New("sorry")
}
