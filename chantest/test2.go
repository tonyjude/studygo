package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 3)

	fmt.Println("len(c)=", len(c), "cap(c)=", cap(c))

	go func() {
		defer fmt.Println("子进程 end")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子进程正在运行，发送元素i=,", i, "len(c)=", len(c), "cap(c)=", cap(c))
		}
		close(c)
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num=", num)
	}

	fmt.Println("主进程 end")
}
