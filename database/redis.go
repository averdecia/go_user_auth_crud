package database

import (
	"github.com/gomodule/redigo/redis"
)

// Redis database global var
var Redis redis.Conn

// InitRedis function to create connection
func InitRedis() {
	pool := &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "portal.local:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	Redis = pool.Get()
}
