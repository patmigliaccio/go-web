package services

// Config : options for Redis Server
type Config struct {
	SvcPort    int
	DbHost     string
	DbPort     int
	DbPassword string
	Db         int
}
