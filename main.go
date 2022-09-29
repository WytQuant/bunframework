package main

import (
	"github.com/WytQuant/bunframework/connectdb"
	"github.com/WytQuant/bunframework/routes"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
	"log"
	"net/http"
)

func main() {
	// Connect to database
	connectdb.Connect()

	// Create a new server
	app := bunrouter.New(bunrouter.Use(reqlog.NewMiddleware()))

	// Set up router
	routes.UserRoute(app)

	// Server listening to port localhost:4000
	log.Fatalln(http.ListenAndServe(":4000", app))
}
