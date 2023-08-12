package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)

	//启动goroutine来处理工作
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	//增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("任务 : %d", post)
	}

	//当所有工作都处理完时关闭通道
	//以便所有goroutine退出
	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			//通道关闭
			fmt.Printf("工作者: %d : 关机\n", worker)
			return
		}

		fmt.Printf("工作者: %d : 开始 %s\n", worker, task)

		//随机
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Microsecond)

		fmt.Printf("工作者： %d : 完成了 %s\n", worker, task)
	}
}
