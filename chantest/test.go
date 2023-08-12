package main

import "fmt"

func main() {

	test := make(chan int)

	go func() {
		defer fmt.Println("goroutine end...")
		fmt.Println("goroutine start...")
		test <- 666
	}()

	num := <-test
	fmt.Println("num=", num)
	fmt.Println("main end")

}
