package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {

	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))

	fmt.Println("---------------")
	v := reflect.ValueOf(x)
	fmt.Println("Kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("type:", v.Type())
	fmt.Println("value:", v.Float())
}

func TestReI(t *testing.T) {
	var num float64 = 1.32
	// 接口类型变量 ---> 反射类型对象
	value := reflect.ValueOf(num)

	// 反射类型对象 ---> 接口类型变量
	convertValue := value.Interface().(float64)
	fmt.Println(convertValue)

	/*
		反射类型对象---> 接口类型变量，理解为“强制转换”
		Golang对类型要求非常严格，类型一定要完全符合
		一个是*float64，一个是float64，错误则引发panic
	*/
	ptr := reflect.ValueOf(&num)
	convertPtr := ptr.Interface().(*float64)
	fmt.Println(convertPtr)
}

func TestReI2(t *testing.T) {

	p1 := Person{
		Name: "王二狗",
		Age:  39,
		Sex:  "男",
	}
	GetMessage(p1)
}

func GetMessage(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name()) // Person
	fmt.Println("get Kind is :", getType.Kind()) // struct

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields are: ", getValue) // {王二狗 39 男}

	// 获取字段
	/*
		step1: 先获取Type对象：reflect.Type
			NumField()
			Field(index)
		step2: 通过Field()获取每一个Field字段
		step3: Interface(),得到对应的Value
	*/
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("字段名称：%s，字段类型：%s，字段数值：%v\n",
			field.Name, field.Type, value)
	}

	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称：%s，方法类型：%v\n", method.Name, method.Type)
	}
}
