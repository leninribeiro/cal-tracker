package main

import (
	"log"
	"net/http"

	"github.com/leninribeiro/cal-tracker/config"
	"github.com/leninribeiro/cal-tracker/data"
	"github.com/leninribeiro/cal-tracker/routes"
	"github.com/urfave/negroni"
)

func main() {
	data.LoadMongo()
	mux := routes.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(mux)

	err := http.ListenAndServe(config.C.ServerPort, n)
	if err != nil {
		log.Fatalf("There was an error with the server: %s. Terminating!", err)
	}
}
