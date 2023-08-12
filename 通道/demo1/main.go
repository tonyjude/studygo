package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	court := make(chan int)

	wg.Add(2)

	//启动两个选手
	go player("小明", court)
	go player("小红", court)

	//发球
	court <- 1

	//等待游戏结束
	wg.Wait()
}

//模拟一个选手在打网球
func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("选手 %s 赢球\n", name)
			return
		}

		n := rand.Intn(100)
		fmt.Printf("随机数 %d\n", n)
		if n%13 == 0 {
			fmt.Printf("选手 %s 输球\n", name)

			close(court)
			return
		}

		//显示击球数，并将击球数加1
		fmt.Printf("选手 %s 击球  %d\n", name, ball)
		ball++

		//将球打向对手
		court <- ball
	}
}
