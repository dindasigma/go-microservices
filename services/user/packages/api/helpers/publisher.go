package helpers

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

type Publisher struct {
	host  string
	port  string
	topic string
	body  []byte
}

func NewPublisher(host, port, topic string, body []byte) *Publisher {
	return &Publisher{
		host:  host,
		port:  port,
		topic: topic,
		body:  body,
	}
}

// connect to DB
func (c *Publisher) Publish() error {
	config := nsq.NewConfig()

	nsqaddress := fmt.Sprintf("%s:%s", c.host, c.port)
	producer, err := nsq.NewProducer(nsqaddress, config)
	if err != nil {
		return err
	}

	err = producer.Publish(c.topic, c.body)
	if err != nil {
		return err
	}

	producer.Stop()

	return nil
}
