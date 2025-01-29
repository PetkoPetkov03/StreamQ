package utils

import (
  "github.com/gin-gonic/gin"
)

type RouteToMethodMap struct {
  Route string
  Method string
}

func SetupRouter() *gin.Engine {
  r := gin.Default()

  r.LoadHTMLGlob("templates/*")
  r.Static("/static", "./static")

  return r
}

func SetUpRoutes() map[RouteToMethodMap]func (*gin.Context) {
  funcHandler := make(map[RouteToMethodMap]func (*gin.Context))

  return funcHandler
}

