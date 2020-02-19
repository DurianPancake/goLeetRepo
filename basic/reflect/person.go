package reflect

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello,", msg)
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名：%s, 年龄: %d, 性别：%s", p.Name, p.Age, p.Sex)
}
