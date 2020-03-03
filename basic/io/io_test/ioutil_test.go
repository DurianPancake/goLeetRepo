package io

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestUtil1(t *testing.T) {

	str := "你好啊！ Hello world"

	reader := strings.NewReader(str)
	data, err := ioutil.ReadAll(reader)
	fmt.Println("err:", err)
	fmt.Println("data:", data)
	fmt.Println(string(data))
}

func TestReadDir(t *testing.T) {

	dir := "D://"
	infos, err := ioutil.ReadDir(dir)
	fmt.Println("err:", err)
	for i, info := range infos {
		fmt.Printf("第%d个文件【%s】是否是目录 %t\n", i+1, info.Name(), info.IsDir())
	}

}
