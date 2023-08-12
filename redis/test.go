package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "47.100.60.204:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()

	/*
		_, err = conn.Do("HSet", "user01", "name", "张三")
		if err != nil {
			fmt.Println("hset err=", err)
			return
		}

		_, err = conn.Do("HSet", "user01", "age", 18)
		if err != nil {
			fmt.Println("hset err=", err)
			return
		}

		name, err := redis.String(conn.Do("HGet", "user01", "name"))
		if err != nil {
			fmt.Println("set err=", err)
			return
		}

		age, err := redis.String(conn.Do("HGet", "user01", "age"))
		if err != nil {
			fmt.Println("set err=", err)
			return
		}
	*/
	/*

		_, err = conn.Do("MSet", "name2", "李四", "address", "上海", "job", "程序员")
		r, err := redis.Strings(conn.Do("MGet", "name2", "address", "job"))
		for _, v := range r {
			fmt.Println(v)
		}

		_, err = conn.Do("HMSet", "user03", "name", "小白龙", "age", 18, "job", "龙族")
		if err != nil {
			fmt.Println("HMSet err=", err)
			return
		}

		r, err = redis.Strings(conn.Do("HMGet", "user03", "name"))
		if err != nil {
			fmt.Println("HMGet err=", err)
			return
		}


			_, err = conn.Do("expire", "name", 10)
			if err != nil {
				fmt.Println("expire err=", err)
				return
			}
	*/

	//fmt.Println(r)
	/*
		_, err = conn.Do("lpush", "heroList", "no1:宋江", 30, "no2:卢俊义", 28)
		if err != nil {
			fmt.Println("lpush err=", err)
			return
		}
	*/

	r, err := redis.String(conn.Do("rpop", "heroList"))
	if err != nil {
		fmt.Println("rpop err=", err)
		return
	}

	fmt.Println(r)
	//fmt.Println("ok", name, age)

}
