package controllers

import (
	"log"

	"github.com/nsqio/go-nsq"
)

var (
	TelegramController telegramControllerInterface = &telegramController{}
)

type telegramControllerInterface interface {
	SendNewUserNotification(*nsq.Message) error
}

type telegramController struct{}

func (c telegramController) SendNewUserNotification(m *nsq.Message) error {
	userid := string(m.Body)
	log.Printf("we are sending telegram notification to admin of new user with id %s", userid)
	// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
	return nil

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	//return errors.New("sorry")
}