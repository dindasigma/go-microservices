package controllers

import (
	"net/http"

	"github.com/dindasigma/go-microservices-user/packages/api/utils/responses"
)

var (
	HomeController homeControllerInterface = &homeController{}
)

type homeControllerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type homeController struct{}

func (c *homeController) Index(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to The Machine")
}
