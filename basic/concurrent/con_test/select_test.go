package con

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	/*
		分支语句： if, switch, select
		select语句类型于switch语句
			但是select语句会随机执行一个可运行的case
			如果没有case可以运行，要看是否有default，如果有就执行default,
				否则就进入阻塞，直到有case可以运行
	*/
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 300
	}()

	select {
	case num1 := <-ch1:
		fmt.Println("ch1中获取的数据。。", num1)
	case num2, ok := <-ch2:
		if ok {
			fmt.Println("ch2中读取的数据。。", num2)
		} else {
			fmt.Println("ch2通道已经关闭..")
		}
	default:
		fmt.Println("default语句")
	}
	fmt.Println("main ... over ...")
}

func TestSelect2(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()

	select {
	case <-ch1:
		fmt.Println("case1可以执行")
	case <-ch2:
		fmt.Println("case2可以执行")
	case <-time.After(3 * time.Second):
		fmt.Println("case3执行，，timeout。。。")
		//default:
		//	fmt.Println("执行了default")
	}
}
