package pkgs

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	/*
		1. func NewTimer(d Duration) *Timer
			创建一个计时器，d时间以后触发
	*/
	timer := time.NewTimer(3 * time.Second)
	fmt.Printf("%T\n", timer)
	fmt.Println(time.Now())

	// 此处等待channel中国年的数值，会阻塞3秒
	ch := timer.C
	fmt.Println(<-ch)
}

func TestTimer2(t *testing.T) {
	// 新建一个计时器
	timer := time.NewTimer(5 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("Timer 2 结束了 。。。 开始 。。。")
	}()
	time.Sleep(3 * time.Second)
	flag := timer.Stop()
	if flag {
		fmt.Println("Timer 2 停止了")
	}
}

func TestTimer3(t *testing.T) {
	/*
		func After(d Duration) <- chan Time
			返回一个通道：chan, 存储的是d时间间隔之后的当前时间
		等价于 return time.NewTimer(d).C
	*/
	ch := time.After(3 * time.Second)
	fmt.Printf("%T\n", ch)
	fmt.Println(time.Now())

	time := <-ch
	fmt.Println(time)
}
