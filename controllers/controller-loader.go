package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type ControllerMethod struct {
	method  string
	route   string
	handler func(*gin.Context)
}

type ControllerInterface interface {
	Get(string, func(ctx *gin.Context))
	Post(string, func(ctx *gin.Context))
	Delete(string, func(ctx *gin.Context))
	Put(string, func(ctx *gin.Context))

	Init(*gin.Engine)
}

type Controller struct {
	entries []ControllerMethod
}

type Handler = func(*gin.Context)

func (controller *Controller) Get(route string, handler Handler) {
	entrie := ControllerMethod{
		method:  "GET",
		route:   route,
		handler: handler,
	}

	controller.entries = append(controller.entries, entrie)
}

func (controller *Controller) Post(route string, handler Handler) {
	entrie := ControllerMethod{
		method:  "POST",
		route:   route,
		handler: handler,
	}

	controller.entries = append(controller.entries, entrie)
}

func (controller *Controller) Delete(route string, handler Handler) {
	entrie := ControllerMethod{
		method:  "DELETE",
		route:   route,
		handler: handler,
	}

	controller.entries = append(controller.entries, entrie)
}

func (controller *Controller) Put(route string, handler Handler) {
	entrie := ControllerMethod{
		method:  "PUT",
		route:   route,
		handler: handler,
	}

	controller.entries = append(controller.entries, entrie)
}

var (
	controller            Controller
	registeredControllers []func(*Controller)
)

func FetchController() Controller {
	if controller.entries == nil {
		controller = Controller{
			entries: make([]ControllerMethod, 0),
		}

		for _, register := range registeredControllers {
			register(&controller)
		}
	}
	return controller
}

func RegisterController(registerFunc func(*Controller)) {
	registeredControllers = append(registeredControllers, registerFunc)
}

func (controller *Controller) Init(router *gin.Engine) {
	for _, v := range controller.entries {
		switch v.method {
		case "GET":
			router.GET(v.route, v.handler)
		case "POST":
			router.POST(v.route, v.handler)
		case "DELETE":
			router.DELETE(v.route, v.handler)
		case "PUT":
			router.PUT(v.route, v.handler)
		default:
			log.Fatalf("Controller initialization issue %v %v", v.route, v.method)
		}
	}
}
