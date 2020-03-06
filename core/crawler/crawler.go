package crawler

import (
	"fmt"
	"go-redis-crawler/core/util"
	"sync"
)

type Crawler struct {
	Redpool util.Redis
	Number  int
	Timeout int
	Retry   int
}

var wg *sync.WaitGroup

func (c *Crawler)Run() {
	wg = &sync.WaitGroup{}
	for i := 0 ; i < c.Number; i ++ {
		wg.Add(1)
		go c.read()
	}
	wg.Wait()
}

func (c *Crawler)read() error{
	defer wg.Done()
	conn := c.Redpool.Pool.Get()
	if err := conn.Err(); err != nil {
		return err
	}
	defer conn.Close()
	count := 0
	for{
		result, e := conn.Do("brpop", "task", c.Timeout)
		if e == nil && result == nil {
			if count == c.Retry {
				break
			}
			count ++
			fmt.Println("retrying %d time(s)", count)
			continue
		}
		if e != nil {
			fmt.Printf("Error occurred while poping from list: %s\n", e)
		}
		fmt.Printf("%s\n",result)
		count = 0
	}
	fmt.Println("thread exited")
	return nil
}