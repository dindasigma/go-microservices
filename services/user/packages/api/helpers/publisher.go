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

func (p Publisher) Publish() error {
	config := nsq.NewConfig()

	nsqaddress := fmt.Sprintf("%s:%s", p.host, p.port)
	producer, err := nsq.NewProducer(nsqaddress, config)
	if err != nil {
		return err
	}

	err = producer.Publish(p.topic, p.body)
	if err != nil {
		return err
	}

	producer.Stop()
	return nil
}
