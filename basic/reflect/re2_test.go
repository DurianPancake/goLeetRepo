package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRe2(t *testing.T) {
	var num float64 = 1.23
	fmt.Println("Num的数值是：", num)

	// 需要操作指针
	// 通过reflect.ValueOf()获取num的Value对象
	ptr := reflect.ValueOf(&num) // 注意参数必须是指针类型才能修改其值
	v := ptr.Elem()

	fmt.Println("类型：", v.Type())         //float64
	fmt.Println("是否可以修改数据：", v.CanSet()) // true

	// 重新赋值
	v.SetFloat(3.14)
	fmt.Println(num)
}

func TestRe3(t *testing.T) {
	// 如果reflect.ValueOf的参数不是指针类型
	var num = 1.23
	value := reflect.ValueOf(num)
	fmt.Println(value.CanSet())
	//value.SetFloat(3.22) //panic: reflect: reflect.flag.mustBeAssignable using unaddressable value
	//value.Elem()	// panic: reflect: call of reflect.Value.Elem on float64 Value
}

func TestRe4(t *testing.T) {

	p := Person{
		Name: "葛二蛋",
		Age:  18,
		Sex:  "男",
	}
	// 通过反射，更改对象的数值：前提也是数值可以被更改
	fmt.Printf("%T\n", p)
	p1 := &p
	fmt.Printf("%T\n", p1)

	// 改变数值
	value := reflect.ValueOf(p1)
	if value.Kind() == reflect.Ptr {
		newValue := value.Elem()
		fmt.Println(newValue.CanSet()) // true

		f1 := newValue.FieldByName("Name")
		f1.SetString("朱一旦")
		f3 := newValue.FieldByName("Sex")
		f3.SetString("非洲猛男")
		fmt.Println(p)
	}
}
