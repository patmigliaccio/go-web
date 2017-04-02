package tune

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/patmigliaccio/go-web/app/services"
	"gopkg.in/gin-gonic/gin.v1"
)

// TuneService : instance of the Tune Service
type TuneService struct {
	Host string
}

// getDB : returns a client instance of a Redis Server
func (ts *TuneService) getDB(cfg services.Config) (*redis.Client, error) {
	addr := cfg.DbHost + ":" + strconv.Itoa(cfg.DbPort)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.DbPassword,
		DB:       cfg.Db,
	})

	pong, err := client.Ping().Result()
	if pong == "PONG" {
		fmt.Printf("[REDIS] Server started on " + addr + "\n")
	}

	return client, err
}

// Run : starts the TuneService
func (ts *TuneService) Run(cfg services.Config) error {
	db, err := ts.getDB(cfg)
	if err != nil {
		return err
	}

	tr := &Resource{Db: db}

	r := gin.New()

	r.GET("/tune/:id", tr.GetTune)
	r.POST("/tune", tr.AddTune)

	ts.Host = cfg.SvcHost + ":" + strconv.Itoa(cfg.SvcPort)

	go r.Run(ts.Host)

	return nil
}
