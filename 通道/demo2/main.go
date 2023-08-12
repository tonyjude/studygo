package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//创建一个无缓冲的通道
	baton := make(chan int)

	//为最后一位跑步者将计数加1
	wg.Add(1)

	//第一位跑步者持有接力棒
	go Runner(baton)

	//开始比赛
	baton <- 1

	//等待比赛结束
	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int
	runner := <-baton

	fmt.Printf("跑步者 %d 带着接力棒跑步\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("跑步者 %d 到达接力地点\n", newRunner)
		go Runner(baton)
	}

	//围绕跑道跑
	time.Sleep(100 * time.Microsecond)

	//比赛结束了吗？
	if runner == 4 {
		fmt.Printf("跑步者 %d 完成, 接力赛结束\n", runner)
		wg.Done()
		return
	}

	//将接力棒交给下一位跑步者
	fmt.Printf("跑步者 %d 进行接力棒交接 %d\n", runner, newRunner)

	baton <- newRunner

}
