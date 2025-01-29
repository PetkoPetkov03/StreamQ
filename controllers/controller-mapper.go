package controllers

import (
  "github.com/gin-gonic/gin"
  "github.com/PetkoPetkov/streamq-backend/utils"
)

func MapController(funcs map[utils.RouteToMethodMap]func (*gin.Context)) {
  var ftm utils.RouteToMethodMap
  ftm.Route = "/register"
  ftm.Method = "GET"

  funcs[ftm] = RegisterView

  ftm.Method = "POST"
  funcs[ftm] = RegisterProcessor

  ftm.Route = "/"
  ftm.Method = "GET"

  funcs[ftm] = IndexView
}
