package document

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/patmigliaccio/go-web/app/services"
	"gopkg.in/gin-gonic/gin.v1"
)

// DocumentService : instance of the Document Service
type DocumentService struct {
}

// getDB : returns a client instance of a Redis Server
func (ds *DocumentService) getDB(cfg services.Config) (*redis.Client, error) {
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

// Run : starts the DocumentService
func (ds *DocumentService) Run(cfg services.Config) error {
	db, err := ds.getDB(cfg)
	if err != nil {
		return err
	}

	dr := &Resource{Db: db}

	r := gin.New()

	r.GET("/document/:id", dr.GetDocument)

	port := ":" + strconv.Itoa(cfg.SvcPort)

	go r.Run(port)

	return nil
}
