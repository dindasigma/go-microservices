package controllers

import (
	"log"
	"net/http"

	"github.com/nsqio/go-nsq"
	"gl.atisicloud.com/dinda/sim-infinyscloud-utils/responses"
)

var (
	HomeController homeControllerInterface = &homeController{}
)

type homeControllerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type homeController struct{}

func (c *homeController) Index(w http.ResponseWriter, r *http.Request) {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("app_nsqd:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	//for i:= 0; i < 1000; i++ {
	//err := producer.Publish("test-nsq", []byte("test" + strconv.Itoa(i)))
	err = producer.Publish("clicks", []byte("haloo"))
	if err != nil {
		log.Fatal(err)
	}
	//}
	producer.Stop()

	responses.JSON(w, http.StatusOK, "Welcome to The Machine")
}
