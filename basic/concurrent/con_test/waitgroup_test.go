package con

import (
	"fmt"
	"sync"
	"testing"
)

// this demo is showing how to use wait_group and how it works
func TestBasic1(t *testing.T) {
	go f1()
	go f2()
	// there are most probably showing nothing before the main goroutine run over
}

var waitGroup sync.WaitGroup

func TestWG(t *testing.T) {

	waitGroup.Add(2) // magic number, todo: how to require task number
	go f1()
	go f2()

	// wait
	fmt.Println("main 进入阻塞状态。。 等待waitGroup中的子Goroutine结束")
	waitGroup.Wait()
	fmt.Println("main 解除阻塞 ....")
}

///////////////////////////////////////
func f1() {
	for i := 1; i < 10; i++ {
		fmt.Println("f1()函数中打印，A", i)
	}
	waitGroup.Done()
}

func f2() {
	defer waitGroup.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("f2()函数中打印，B", i)
	}
}
