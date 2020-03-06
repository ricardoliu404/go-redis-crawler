package test

import (
	"go-redis-crawler/core/util"
	"testing"
)

func TestRedis(t *testing.T) {
	r := util.GetRedis("localhost:6379")
	conn := r.Pool.Get()
	if err := conn.Err(); err != nil {
		t.Errorf("Something wrong get connection with redis: %s", err)
	}
	defer conn.Close()
	_, err := conn.Do("lpush", "task", "string")
	if err != nil {
		t.Errorf("Error while writing to redis: %s", err)
	}
}