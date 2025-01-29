package main

import (
	"log"

	"github.com/PetkoPetkov/streamq-backend/controllers"
	"github.com/PetkoPetkov/streamq-backend/streamqsql/schemas"
	"github.com/PetkoPetkov/streamq-backend/utils"
)

func main() {
	queries, err := schemas.SetUpDBConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	r := utils.SetupRouter()

	schemas.SetQueryCaller(queries)

	if err != nil {
		log.Fatal(err)
	}

	funcs := utils.SetUpRoutes()

	controllers.MapController(funcs)

	for k, v := range funcs {
		if k.Method == "GET" {
			r.GET(k.Route, v)
		}

		if k.Method == "POST" {
			r.POST(k.Route, v)
		}
	}

	r.Run()
}
