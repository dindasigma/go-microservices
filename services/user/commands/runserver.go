package main

import (
	"log"

	"github.com/dindasigma/go-microservices-user/packages/api"
)

// @title User Microservices
// @version 1.0
// @description Handled users management and login.
// @termsOfService http://swagger.io/terms/

// @contact.name Dinda
// @contact.email dindasigma@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := api.Run(); err != nil {
		log.Fatal(err)
	}

}
