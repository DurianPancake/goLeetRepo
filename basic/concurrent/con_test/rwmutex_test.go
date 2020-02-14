package con

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var lock *sync.RWMutex

func TestRwLockRead(t *testing.T) {

	lock = new(sync.RWMutex)
	wg := new(sync.WaitGroup)

	wg.Add(2)
	// 多个读操作可以同时进行
	go readData(1, wg)
	go readData(2, wg)
	wg.Wait()
	/*
		2 开始读： read start
		2 正在读取数据： reading
		1 开始读： read start
		1 正在读取数据： reading
		2 读结束：read over...
		1 读结束：read over...
	*/
}

func TestRWLockWrite(t *testing.T) {

	lock = new(sync.RWMutex)
	wg := new(sync.WaitGroup)

	wg.Add(2)
	// 多个写操作不能同时进行
	go writeData(1, wg)
	go writeData(2, wg)
	wg.Wait()
	/*
		2 开始写 write start ..
		2 正在写 writing ...
		2 写结束：write over ...
		1 开始写 write start ..
		1 正在写 writing ...
		1 写结束：write over ...
	*/
}

func TestMixed(t *testing.T) {
	lock = new(sync.RWMutex)
	wg := new(sync.WaitGroup)

	// 写时无法进行读操作
	wg.Add(4)
	go readData(1, wg)
	go writeData(1, wg)
	go readData(2, wg)
	go writeData(2, wg)
	wg.Wait()
	/*
		2 开始写 write start ..
		2 正在写 writing ...
		2 写结束：write over ...
		2 开始读： read start
		2 正在读取数据： reading
		1 开始写 write start ..
		1 开始读： read start
		2 读结束：read over...
		1 正在写 writing ...
		1 写结束：write over ...
		1 正在读取数据： reading
		1 读结束：read over...
	*/
}

func readData(i int, wg *sync.WaitGroup) {
	fmt.Println(i, "开始读： read start")
	defer wg.Done()

	lock.RLock() // 读锁
	fmt.Println(i, "正在读取数据： reading")
	time.Sleep(1 * time.Second)
	lock.RUnlock() // 读操作解锁
	fmt.Println(i, "读结束：read over...")
}

func writeData(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i, "开始写 write start ..")
	lock.Lock() // 写操作上锁
	fmt.Println(i, "正在写 writing ...")
	lock.Unlock() // 写操作解锁
	fmt.Println(i, "写结束：write over ...")
	/*

	 */
}
