package test

import (
	"go-redis-crawler/core/crawler"
	"go-redis-crawler/core/util"
	"testing"
)

func TestCrawler(t *testing.T) {
	c := &crawler.Crawler{
		Redpool: util.GetRedis("localhost:6379"),
		Number:  3,
		Timeout: 3,
		Retry:   3,
	}
	c.Run()
}