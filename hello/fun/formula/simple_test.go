package formula

import (
	"fmt"
	"goSecond/hello/fun/formula/model"
	"goSecond/hello/fun/formula/utils"
	"testing"
)

var (
	f1 = "A=3 && B!=4 ||  C>5"
	f2 = "A=3 || B!=4 && C>5"
	f3 = "A=3 &&( B!=4 ||  C>5)"
	f4 = "A7YWFGS=9||JFKGL>4|| 胥志龙=男&&(女朋友=无)"
	f5 = "A7YWFGS=9||((JFKGL>4|| 胥志龙=男)&&(女朋友=无))"
)

func TestUnix(t *testing.T) {
	s := "sdfa"
	operate := model.MatchOperate(s, 2)
	fmt.Println(operate == model.Null)
}

func TestGen(t *testing.T) {
	condition, _ := utils.GenerateCondition(f5)
	fmt.Println("condition:", condition)
}
