package main

import "fmt"

func f(ch chan int) {
	ret := <-ch
	fmt.Println(ret)
}

func main() {
	var ch chan int
	//ch = make(chan int, 100)

	ch = make(chan int)
	go f(ch)

	ch <- 100

	fmt.Println("hello")

}
