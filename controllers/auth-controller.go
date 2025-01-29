package controllers

import (
	"net/http"

	"github.com/PetkoPetkov/streamq-backend/services"
	"github.com/gin-gonic/gin"
)

func RegisterView(ctx *gin.Context) {
  ctx.HTML(http.StatusOK, "register.html", gin.H{
  })
}

func RegisterProcessor(ctx *gin.Context) {
  req := services.UserAuthReq{} 

  if err := ctx.ShouldBindJSON(&req); err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })

    return
  }

  err := services.GetAuthService().Register(ctx, req)

  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H {
      "err": err.Error(),
    })

    return
  }

  ctx.JSON(http.StatusOK, nil)
}
