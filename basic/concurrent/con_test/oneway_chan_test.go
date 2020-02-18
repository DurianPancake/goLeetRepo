package con

import (
	"fmt"
	"testing"
)

func TestOneWayChan(t *testing.T) {
	/*
		单向：定向
		chan <- T, 只支持写
		<- chan T, 只读
	*/
	ch1 := make(chan int) // 双向，读，写
	//ch2 := make(chan <- int)	// 单向，只能写
	//ch3 := make(<- chan int)	// 单向，只能读

	go s1(ch1) //
	//s1(ch2) //

	data := <-ch1
	fmt.Println("func1函数中写出的数据是：", data)
}

// one function which can only send
func s1(ch chan<- int) {
	// 函数内部，用作权限限制，安全保护
	ch <- 100
	fmt.Println("fun1 函数结束...")
}
