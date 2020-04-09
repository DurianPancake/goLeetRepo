package io

import (
	"fmt"
	"os"
	"testing"
)

func TestEdit(t *testing.T) {

	filePath := "c://Users/LLH/Desktop/报表子系统0401/form/1214389265371369472.json"
	file, _ := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE, 0777)
	fmt.Println(file)
	// 清空文件
	defer file.Close()

	//ioutil.ReadFile(filePath)
}
