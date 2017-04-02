package services

// Config : options for Redis Server
type Config struct {
	SvcHost     string
	SvcPort     int
	SvcProtocol string
	DbHost      string
	DbPort      int
	DbPassword  string
	Db          int
}
