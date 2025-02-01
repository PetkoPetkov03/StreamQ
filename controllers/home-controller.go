package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(controller *Controller) {

	controller.Get("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
}

func init() {
	RegisterController(HomeController)
}
