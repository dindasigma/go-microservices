package application

import (
	"fmt"
	"log"
	"sync"

	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	initializeConsumers()

	wg.Wait()
}
