package main
 
import (
	"context"
	"fmt"
	"time"
)

func doTask(ctx context.Context, name string) {
	
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s task canceled\n", name)
			return
		default:
			fmt.Printf("%s task is running\n", name)
			time.Sleep(1 * time.Second)
		}
	}

}

 
func main() {
	//创建顶层Context
	ctx := context.Background()
	
	//派生2个子Context, 并将它们传递给2个任务
	ctx1, cancel := context.WithCancel(ctx)
	ctx2, cance2 := context.WithCancel(ctx)
	
	go doTask(ctx1, "Task 1")
	go doTask(ctx2, "Task 2")
	
	time.Sleep(5 * time.Second)
	cancel()
	
	time.Sleep(5 * time.Second)
	cance2()
	
	time.Sleep(2 * time.Second)
}
 