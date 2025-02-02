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

	controller := controllers.FetchController()

	controller.Init(r)

	r.Run()
}
