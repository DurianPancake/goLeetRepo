package formula

import (
	"fmt"
	"goSecond/hello/fun/formula/utils"
	"testing"
	"time"
)

var (
	f1 = "A=3"
	f2 = "A=3 || B!=4 && (C>5)"
	f3 = "A=3 &&( B!=4 ||  C>5)"
	f4 = "A7YWFGS=9||JFKGL>4|| 胥志龙=男&&(女朋友=无)"
	f5 = "A7YWFGS>9||((JFKGL>=4|| 胥志龙like男)&&(女朋友=无))"
)

func TestGen(t *testing.T) {
	condition, _ := utils.GenerateCondition(f5)
	fmt.Println("condition:", condition)
}

func TestGen2(t *testing.T) {
	condition, _ := utils.GenerateCondition(f1)
	fmt.Println("condition:", condition)
	condition, _ = utils.GenerateCondition(f2)
	fmt.Println("condition:", condition)
	condition, _ = utils.GenerateCondition(f3)
	fmt.Println("condition:", condition)
	condition, _ = utils.GenerateCondition(f4)
	fmt.Println("condition:", condition)
	condition, _ = utils.GenerateCondition(f5)
	fmt.Println("condition:", condition)
}

func TestTime(t *testing.T) {
	start := time.Now().UnixNano()
	for i := 0; i < 100000; i++ {
		_, _ = utils.GenerateCondition(f5)
	}
	fmt.Println("耗时：", (time.Now().UnixNano()-start)/1000000, "ms")
	//fmt.Println("condition:", condition)
}
