package main

import (
	"fmt"
)

func push(numChan chan int, len int) {
	for i := 1; i <= len; i++ {
		numChan <- i
	}
	close(numChan)
}

func count(numChan chan int, addChan chan int, exitChan chan bool) {
	for {
		tempNum, ok := <-numChan
		if !ok {
			break
		}

		res := 0
		for i := 1; i <= tempNum; i++ {
			res += i
		}

		addChan <- res
	}

	exitChan <- true
	close(addChan)
}

func main() {

	len := 10
	numChan := make(chan int, len)
	addChan := make(chan int, len)
	exitChan := make(chan bool, 1)

	go push(numChan, len)

	for i := 0; i < 8; i++ {
		go count(numChan, addChan, exitChan)
	}

	for {
		ret, ok := <-addChan
		if !ok {
			break
		}

		fmt.Printf("ret=%d\n", ret)
	}

	//<-exitChan
}
