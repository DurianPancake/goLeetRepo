package pkgs

import (
	"fmt"
	"runtime"
	"testing"
)

func TestRuntime1(t *testing.T) {
	// 获取goroot目录
	fmt.Println("GOROOT -->", runtime.GOROOT())
	// 获取操作系统
	fmt.Println("os/platform", runtime.GOOS)

	// 获取逻辑CPU的数量
	fmt.Println("逻辑CPU的数量--->", runtime.NumCPU())

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine..")
		}
	}()

	for i := 0; i < 4; i++ {
		// 让出时间片，先让别的goroutine执行
		runtime.Gosched()
		fmt.Println("main...")
	}
}
