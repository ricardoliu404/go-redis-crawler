package test

import (
	"go-redis-crawler/core/headman"
	"go-redis-crawler/core/util"
	"testing"
	"time"
)

func TestHeadman(t *testing.T) {
	h := &headman.Headman{Redpool:util.GetRedis("localhost:6379")}
	str := []string {
		"我是",
		"一个",
		"兵",
		"来自",
		"老白",
		"姓",
		"呢",
		"恩",
	}
	for i := 0 ; i < 10 ; i ++{
		h.BatchAppend(str)
		time.Sleep(5 * time.Second)
	}
}