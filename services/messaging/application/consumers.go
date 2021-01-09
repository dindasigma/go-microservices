package application

import (
	"fmt"
	"log"
	"os"

	"github.com/dindasigma/go-microservices-messaging/controllers"
	"github.com/nsqio/go-nsq"
)

func initializeConsumers() {
	config := nsq.NewConfig()

	newConsumer(config, "new_user", "email_welcome", controllers.EmailController.SendWelcome)
	newConsumer(config, "new_user", "telegram_notification", controllers.TelegramController.SendNewUserNotification)
}

func newConsumer(config *nsq.Config, topic string, channel string, handler nsq.HandlerFunc) {
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}

	// set the Handler for messages received by this Consumer.
	consumer.AddHandler(nsq.HandlerFunc(handler))

	nsqaddress := fmt.Sprintf("%s:%s", os.Getenv("NSQD_SERVICE_HOST"), os.Getenv("NSQD_SERVICE_PORT"))
	err = consumer.ConnectToNSQD(nsqaddress)
	if err != nil {
		log.Fatal(err)
	}
}
