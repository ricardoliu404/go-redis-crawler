package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

var (
	explore = []string{
		"Chrome",
		"Firefox",
		"Internet+Explorer",
		"Safari",
	}
	urlpattern = "http://useragentstring.com/pages/useragentstring.php?name=%s"
	reg = "(Mozilla/.*?)</a>"
	wg *sync.WaitGroup
	defaultUA = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36"
)

func GetUA() {
	wg = &sync.WaitGroup{}
	r := GetRedis("localhost:6379")

	for _, x := range explore {
		wg.Add(1)
		go getBrowser(x, r)
	}
	wg.Wait()
}

func getBrowser(xs string, redis Redis) {
	conn := r.Pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		fmt.Printf("Something wrong get connection with redis: %s", err)
	}
	//fmt.Println(fmt.Sprintf(urlpattern, xs))
	defer wg.Done()
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", fmt.Sprintf(urlpattern, xs), nil)
	reqest.Header.Add("User-Agent", defaultUA)
	res, _ := client.Do(reqest)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	ua := regexp.MustCompile(reg)
	matchUA := ua.FindAllStringSubmatch(string(body), -1)

	args := []interface{}{"userAgent"}

	for _, v := range matchUA {
		args = append(args, v[1])
		fmt.Println(v[1])
	}
	//fmt.Printf("args: len: %d\n", len(args))
	_, err = conn.Do("sadd", args...)
	if err != nil {
		fmt.Printf("Error while writing useragent to redis: %s", err)
	}
}