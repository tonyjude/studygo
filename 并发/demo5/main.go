package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex   //互斥锁
	rwLock sync.RWMutex //读写锁 读的次数远远大于写的次数
)

func read() {
	defer wg.Done()
	rwLock.RLock()
	//lock.Lock()
	//fmt.Println(x)
	time.Sleep(time.Microsecond * 1)
	rwLock.RUnlock()
	//lock.Unlock()
}

func write() {
	defer wg.Done()

	rwLock.Lock()
	//lock.Lock()
	x = x + 1
	time.Sleep(time.Microsecond * 10)
	rwLock.Unlock()
	//lock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Printf("耗费了%v.", end.Sub(start))
}
