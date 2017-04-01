package main

import (
	"net/http"

	"fmt"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/patmigliaccio/go-web/server/api"
	"github.com/spf13/viper"
) // TODO: Move local mports to common cmd directory

// LoadConfig : reads in 'environment.toml' to Viper config
func LoadConfig() {
	viper.SetConfigName("environment")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config file: %s", err))
	}
}

// LoadAPI : configures api service
func LoadAPI(r *httprouter.Router) {
	apiRoot := viper.GetString("web.api.root")
	api.AddRoutes(apiRoot, r)
}

// LoadAssets : adds asset folder configuration
func LoadAssets(r *httprouter.Router) {
	assets := viper.GetStringMapString("web.assets")
	route, location := assets["route"]+"/*filepath", assets["location"]
	r.ServeFiles(route, http.Dir(location))
}

// Index : serves the initial app
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fs := http.FileServer(http.Dir("../public"))
	fs.ServeHTTP(w, r)
}

func main() {
	LoadConfig()

	r := httprouter.New()
	r.GET("/", Index)

	LoadAPI(r)
	LoadAssets(r)

	port := ":" + strconv.Itoa(viper.GetInt("server.port"))

	err := http.ListenAndServe(port, r)
	if err != nil {
		panic(fmt.Errorf("server error: %s", err))
	}
}
