package publisher

import (
	"fmt"
	"log"
	"os"

	"github.com/nsqio/go-nsq"
)

func Publish(topic string, body []byte) {
	config := nsq.NewConfig()

	nsqaddress := fmt.Sprintf("%s:%s", os.Getenv("NSQD_SERVICE_HOST"), os.Getenv("NSQD_SERVICE_PORT"))
	producer, err := nsq.NewProducer(nsqaddress, config)
	if err != nil {
		log.Fatal(err)
	}

	err = producer.Publish(topic, body)
	if err != nil {
		log.Fatal(err)
	}

	producer.Stop()
}
