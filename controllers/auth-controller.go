package controllers

import (
	"net/http"

	"github.com/PetkoPetkov/streamq-backend/services"
	"github.com/gin-gonic/gin"
)

func AuthController(controller *Controller) {

	controller.Get("/register", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "register.html", gin.H{})
	})

	controller.Post("/register", func(ctx *gin.Context) {
		req := services.UserAuthReq{}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "request error",
			})

			return
		}

		err := services.GetAuthService().Register(ctx, req)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "service error",
			})

			return
		}

		ctx.JSON(http.StatusOK, nil)
	})

	controller.Get("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})

	controller.Post("/login", func(ctx *gin.Context) {
		req := services.UserAuth{}

		session, err := services.GetAuthService().Login(ctx, req)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"session": session,
		})
	})
}

func init() {
	RegisterController(AuthController)
}
