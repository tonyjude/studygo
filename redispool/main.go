package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {

	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "47.100.60.204:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()

	r, err := conn.Do("Set", "name3", "tom猫")
	if err != nil {
		fmt.Println("conn.Do error=", err)
		return
	}

	pool.Close()
	fmt.Println("r=", r)

}
