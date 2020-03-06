package headman

import (
	"fmt"
	"go-redis-crawler/core/util"
)

type Headman struct {
	Redpool util.Redis
}

func (h *Headman) BatchAppend (list []string) error{
	conn := h.Redpool.Pool.Get()
	if err := conn.Err(); err != nil {
		return err
	}
	defer conn.Close()

	args := []interface{}{"task"}
	for _, val := range list {
		args = append(args, val)
	}

	result, e := conn.Do("lpush", args...)
	if e != nil {
		return e
	}
	fmt.Println(result)
	return nil
}