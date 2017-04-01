package main

import (
	"net/http"

	"fmt"
	"strconv"

	"github.com/patmigliaccio/go-web/server/api"
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
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
func LoadAPI(r *gin.Engine) {
	apiRoot := viper.GetString("web.api.root")
	api.AddRoutes(apiRoot, r)
}

// LoadAssets : adds asset folder configuration
func LoadAssets(r *gin.Engine) {
	assets := viper.GetStringMapString("web.assets")
	route, location := assets["route"], assets["location"]
	r.StaticFS(route, http.Dir(location))
}

func main() {
	LoadConfig()

	r := gin.Default()
	r.StaticFile("/", "../public")

	LoadAPI(r)
	LoadAssets(r)

	port := ":" + strconv.Itoa(viper.GetInt("server.port"))

	r.Run(port)
}
