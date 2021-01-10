package application

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/dindasigma/go-microservices-messaging/datasources"
	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	datasources.InitializePostgres(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	wg := &sync.WaitGroup{}
	wg.Add(1)

	initializeConsumers()

	wg.Wait()
}
