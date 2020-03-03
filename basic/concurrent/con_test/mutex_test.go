package con

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var ticket = 10

var mutex sync.Mutex // 创建锁头
var wg sync.WaitGroup

func TestBasic(t *testing.T) {
	/*
		4个goroutine，模拟售票口
	*/
	wg.Add(4)
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口 3")
	go saleTickets("售票口4")

	wg.Wait()
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	defer mutex.Unlock()
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出", ticket)
			ticket--
		} else {
			fmt.Println("售罄，没有票了")
			break
		}
		mutex.Unlock()
	}
}
