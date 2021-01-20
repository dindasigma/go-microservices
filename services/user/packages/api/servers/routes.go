package servers

import (
	_ "github.com/dindasigma/go-microservices-user/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/dindasigma/go-microservices-user/packages/api/controllers"
	"github.com/dindasigma/go-microservices-user/packages/api/middlewares"
	httpSwagger "github.com/swaggo/http-swagger"
)

func initializeRoutes() {
	// Home Route
	router.HandleFunc("/", middlewares.SetMiddlewareJSON(controllers.HomeController.Index)).Methods("GET")

	// Login Route
	router.HandleFunc("/login", middlewares.SetMiddlewareJSON(controllers.LoginController.Login)).Methods("POST")
	router.HandleFunc("/authenticated", controllers.LoginController.CheckAuth).Methods("GET")

	// Users routes
	router.HandleFunc("/user", middlewares.SetMiddlewareJSON(controllers.UserController.Get)).Methods("GET")
	router.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(middlewares.SetMiddlewareJSON(controllers.UserController.GetByID))).Methods("GET")
	router.HandleFunc("/user", middlewares.SetMiddlewareJSON(controllers.UserController.Create)).Methods("POST")
	router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(controllers.UserController.Update))).Methods("PUT")
	router.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(controllers.UserController.Delete)).Methods("DELETE")

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}