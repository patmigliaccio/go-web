package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/patmigliaccio/go-web/server/rest"
	"github.com/patmigliaccio/go-web/server/services"
	"github.com/patmigliaccio/go-web/server/services/tune"
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

// LoadConfig : reads in 'environment.toml' to Viper config
func LoadConfig() {
	viper.SetConfigName("environment")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config file: %s", err))
	}
}

// LoadRESTAPI : configures api service
func LoadRESTAPI(r *gin.Engine) {
	apiRoot := viper.GetString("web.rest.root")
	rest.AddRoutes(apiRoot, r)
}

// LoadAssets : adds asset folder configuration
func LoadAssets(r *gin.Engine) {
	assets := viper.GetStringMapString("web.assets")
	route, location := assets["route"], assets["location"]
	r.StaticFS(route, http.Dir(location))
}

// StartServices : runs all services
func StartServices() {
	var redisCfg services.Config

	err := viper.UnmarshalKey("redis", &redisCfg)
	if err != nil {
		panic(fmt.Errorf("error mapping Redis config: %s", err))
	}

	tsvc := tune.TuneService{}
	if err := tsvc.Run(redisCfg); err != nil {
		panic(fmt.Errorf("error starting Tune Service: %s", err))
	}
}

func main() {
	LoadConfig()

	r := gin.New()

	public := viper.GetString("web.location")
	r.StaticFile("/", public)

	LoadRESTAPI(r)
	LoadAssets(r)
	StartServices()

	addr := viper.GetString("server.host") + ":" + strconv.Itoa(viper.GetInt("server.port"))

	if err := r.Run(addr); err != nil {
		panic(fmt.Errorf("error starting app: %s", err))
	}
}
