package util

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type Redis struct {
	Pool *redis.Pool
}

var r *Redis

func init() {
	fmt.Println()
}

func GetRedis(redisUrl string) Redis {
	if r == nil {
		r = &Redis{
			Pool: &redis.Pool{
				Dial:         func() (redis.Conn, error) {
					return redis.Dial("tcp", redisUrl)
				},
				MaxIdle:      256,
				MaxActive:    0,
				IdleTimeout:  time.Duration(120),
			},
		}
	}
	return *r
}