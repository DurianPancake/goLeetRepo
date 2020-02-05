package io

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

const (
	srcFile = "e://abc/guliang.jpg"
)

func TestRenew(t *testing.T) {
	destFile := srcFile[strings.LastIndex(srcFile, "/")+1:]
	fmt.Println(destFile)
	tempFile := destFile + ".tmp"
	fmt.Println(tempFile)

	source, err := os.Open(srcFile)
	handleErr(err)
	target, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	handleErr(err)
	temp, err := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	handleErr(err)

	defer source.Close()
	defer target.Close()

	// step1 读取临时文件的数据，再Seek
	temp.Seek(0, io.SeekStart)
	bs := make([]byte, 100, 100)
	read, err := temp.Read(bs)
	//handleErr(err)
	countStr := string(bs[:read])
	count, err := strconv.ParseInt(countStr, 10, 32)
	fmt.Println(count)

	// step2 根据count值设置偏移量，设置读和写的位置
	_, _ = source.Seek(count, io.SeekStart)
	_, _ = target.Seek(count, io.SeekStart)
	buffer := make([]byte, 1024, 1024)

	rdNum := -1
	wtNum := -1
	total := int(count)

	// step3 复制文件
	for {
		rdNum, err = source.Read(buffer)
		if err == io.EOF || rdNum == 0 {
			fmt.Println("文件复制完毕", total)
			temp.Close()
			os.Remove(tempFile)
			break
		}
		wtNum, err = target.Write(buffer[:rdNum])
		total += wtNum

		// 将复制的总量，存储到临时文件中
		temp.Seek(0, io.SeekStart)
		temp.WriteString(strconv.Itoa(total))
		//
		//if total > 230000 {
		//	fmt.Printf("total:%d\n",total)
		//	panic("假装断电了！")
		//}
	}

}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
