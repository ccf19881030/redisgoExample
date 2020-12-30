package common

// RedisConnOpt connect redis options
type RedisConnOpt struct {
	Enable   bool
	Host     string
	Port     int32
	Password string
	Index    int32
	TTL      int32
}