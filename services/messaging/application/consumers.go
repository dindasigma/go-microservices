package application

import (
	"fmt"
	"log"
	"os"

	"github.com/dindasigma/go-microservices-messaging/controllers"
	"github.com/nsqio/go-nsq"
)

type Consumer struct {
	host    string
	port    string
	topic   string
	channel string
	handler nsq.HandlerFunc
}

func initializeConsumers() {
	config := nsq.NewConfig()

	email_welcome := &Consumer{
		os.Getenv("NSQD_SERVICE_HOST"),
		os.Getenv("NSQD_SERVICE_PORT"),
		"new_user",
		"email_welcome",
		controllers.EmailController.SendWelcome,
	}
	createConsumer(config, email_welcome)

	telegram_notification := &Consumer{
		os.Getenv("NSQD_SERVICE_HOST"),
		os.Getenv("NSQD_SERVICE_PORT"),
		"new_user",
		"telegram_notification",
		controllers.TelegramController.SendNewUserNotification,
	}
	createConsumer(config, telegram_notification)
}

func createConsumer(config *nsq.Config, c *Consumer) {

	consumer, err := nsq.NewConsumer(c.topic, c.channel, config)
	if err != nil {
		log.Fatal(err)
	}

	// set the Handler for messages received by this Consumer.
	consumer.AddHandler(nsq.HandlerFunc(c.handler))

	nsqaddress := fmt.Sprintf("%s:%s", c.host, c.port)
	err = consumer.ConnectToNSQD(nsqaddress)
	if err != nil {
		log.Fatal(err)
	}
}
