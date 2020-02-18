package con

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestChan(t *testing.T) {

	ch1 := make(chan int)

	go sendData(ch1)
	// for range
	for v := range ch1 { // v <= ch1 阻塞的
		fmt.Println("读取数据：", v)
	}
	fmt.Println("main ... over")
}

func sendData(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		ch <- i
	}
	close(ch)
}

func TestBufChan(t *testing.T) {
	/*
		非缓冲通道：make(chan T)
			一次发送，一次接收，都是阻塞的
		缓冲通道：make(chan T, capacity)
			发送：缓冲区的数据慢了，才会阻塞
			接收：缓冲区的数据空了，才会阻塞
	*/
	ch1 := make(chan int)
	fmt.Println(len(ch1), cap(ch1)) // 0, 0
	// ch1 <- 100 //阻塞式的，需要有其他的goroutine解除阻塞，否则deadlock

	ch2 := make(chan int, 5)
	fmt.Println(len(ch2), cap(ch2)) // 0, 5

	ch2 <- 1
	fmt.Println(len(ch2), cap(ch2)) // 1, 5
	ch2 <- 2
	ch2 <- 3
	ch2 <- 4
	ch2 <- 5
	fmt.Println(len(ch2), cap(ch2)) // 5, 5
	//ch2 <- 6 // 阻塞！fatal error: all goroutines are asleep - deadlock!
}

func TestBufChan2(t *testing.T) {
	ch := make(chan string, 4)

	go sendStr(ch)

	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("读完了。。", ok)
			break
		}
		fmt.Println("\t读取的数据：", v)
	}
	fmt.Println("main. .. over.")
}

func sendStr(ch chan string) {

	for i := 0; i < 10; i++ {
		ch <- "数据" + strconv.Itoa(i)
		fmt.Printf("子goroutine中写出第%d个数据\n", i)
	}
	close(ch)
}
