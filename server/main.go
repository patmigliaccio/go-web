package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/patmigliaccio/go-web/server/api"
) // TODO: Move local mports to common cmd directory

// PORT : constant number to serve on
// TODO: Move to environment config file
const PORT string = "8080"

// Index : serves the initial app
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fs := http.FileServer(http.Dir("../public"))
	fs.ServeHTTP(w, r)
}

func main() {
	// config, err := env.LoadConfig("config.json")

	router := httprouter.New()
	router.GET("/", Index)
	router.ServeFiles("/assets/*filepath", http.Dir("../public/assets"))
	api.AddRoutes(router)

	err := http.ListenAndServe(":"+PORT, router)
	if err != nil {
		log.Fatalln(err)
	}
}
