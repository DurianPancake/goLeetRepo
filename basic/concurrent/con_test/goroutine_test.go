package con_test

import (
	"fmt"
	"testing"
)

// 一个Goroutine打印数字，另外一个goroutine打印字母，观察运行结果
func TestSaint(t *testing.T) {

	// 1.先创建并启动子goroutine, 执行printNum()
	go printNum()

	// 2.main中打印字母
	for i := 1; i <= 100; i++ {
		fmt.Printf("\t主goroutine中打印字母： %c \n", rune(65+i%26))
	}
}

func printNum() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("子goroutine中打印数字： %d\n", i)
	}
}
