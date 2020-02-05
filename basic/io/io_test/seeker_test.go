package io_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

const (
	fileName = "E://abc/a.txt"
)

func TestSeeker(t *testing.T) {
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 读写
	bs := make([]byte, 3, 3)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(4, io.SeekStart)
	file.Read(bs)
	fmt.Println(string(bs))

	seek, _ := file.Seek(2, io.SeekCurrent)
	fmt.Println(seek)
	file.Read(bs)
	fmt.Println(string(bs))
}
