package reflect

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestInvoke(t *testing.T) {
	/*
		通过反射来进行方法的调用
		思路：
		step1:接口变量-->对象反射对象：Value
		step2:获取对应的方法对象：MethodByName()
		step3:将方法对象进行调用
	*/
	person := Person{
		Name: "华农",
		Age:  30,
		Sex:  "2333",
	}

	value := reflect.ValueOf(person)
	fmt.Printf("Kind: %s, type:%s\n", value.Kind(), value.Type())

	method1 := value.MethodByName("PrintInfo")
	fmt.Printf("Kind:%s, type:%s\n", method1.Kind(), method1.Type())

	method1.Call(nil)

	method2 := value.MethodByName("Say")
	fmt.Printf("kind:%s, type:%s\n", method2.Kind(), method2.Type())
	args2 := []reflect.Value{reflect.ValueOf("反射机制")}
	method2.Call(args2)
}

func TestInvoke2(t *testing.T) {
	// 函数的反射
	/*
		思路：函数也是看作接口变量类型
		step1:函数-->反射对象，Value
		step2:kind --> func
		step3:call()
	*/
	f1 := fun1
	value := reflect.ValueOf(f1)
	fmt.Printf("kind:%s, type:%s\n", value.Kind(), value.Type()) // kind:func, type:func()

	value2 := reflect.ValueOf(fun2)
	value3 := reflect.ValueOf(fun3)
	fmt.Printf("kind:%s, type:%s\n", value2.Kind(), value2.Type()) // kind:func, type:func(int, string)
	fmt.Printf("kind:%s, type:%s\n", value3.Kind(), value3.Type()) // kind:func, type:func(int, string) string

	// 通过反射调用函数
	value.Call(nil)
	value2.Call([]reflect.Value{reflect.ValueOf(3), reflect.ValueOf("尼尔")})
	result := value3.Call([]reflect.Value{reflect.ValueOf(300), reflect.ValueOf("塞尔达")})
	fmt.Printf("%T\n", result)
	fmt.Println(len(result))
	fmt.Printf("kind:%s, type:%s\n", result[0].Kind(), result[0].Type())

	s := result[0].Interface().(string)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
}

func fun1() {
	fmt.Println("我是函数fun1()，无参的。。")
}

func fun2(i int, s string) {
	fmt.Println("我是函数fun2(), 有参")
}

func fun3(i int, s string) string {
	fmt.Println("我是函数fun3()，有参且有返回值")
	return s + strconv.Itoa(i)
}
