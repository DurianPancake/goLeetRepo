package io

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestPrint0(t *testing.T) {
	fmt.Println("输入整数n表示n阶")
	numLevel := getStdInNum()
	// 打印三角
	start := time.Now()
	printTriangle(numLevel)
	end := time.Now()
	duration := end.UnixNano() - start.UnixNano()
	fmt.Printf("用时：%f ms", float64(duration/1e6))
}

// 主打印方法
func printTriangle(level int) {
	if level < 0 {
		fmt.Println("输入格式错误")
		return
	}
	// 计算最长数的长度
	maxNumLen := computeMaxLen(level)
	// 计算空格长度
	spaceLen := computeSpaceLen(maxNumLen)
	// 主打印方法
	mainPrint(spaceLen, level)
}

func mainPrint(spaceLen int, n int) {

	filePath := "d://Yanghui.txt"
	file, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	// 清空文件
	_ = file.Truncate(0)
	defer file.Close()
	writer := bufio.NewWriter(file)

	list := make([]*big.Int, 0, n)
	prefixSLen := spaceLen/2 + 1
	for i := 0; i < n; i++ {
		// 写行前空格
		printSpace(writer, prefixSLen, n-(i+1))
		// 计算第i行的元素
		getListByGrade(i, &list)
		// 打印第i行
		printList(writer, &list, spaceLen)
	}
}

func printList(writer *bufio.Writer, list *[]*big.Int, spaceLen int) {

	buffer := bytes.NewBufferString("")
	// 遍历列表
	for i := 0; i < len(*list); i++ {
		value := (*list)[i]
		buffer.WriteString(value.String())
		if i != len(*list)-1 {
			for j := 0; j < spaceLen-len(value.String())+1; j++ {
				buffer.WriteString(" ")
			}
		}
	}
	buffer.WriteString("\r\n")
	buffer.WriteString("\r\n")
	_, _ = writer.WriteString(buffer.String())
	_ = writer.Flush()
}

func getListByGrade(i int, list *[]*big.Int) {
	// 最简情况
	if i == 0 {
		*list = append(*list, big.NewInt(1))
		return
	}
	length := len(*list)
	copyData := make([]*big.Int, length)
	for j := 0; j < length; j++ {
		copyData[j] = big.NewInt((*list)[j].Int64())
	}
	for j := 1; j < i; j++ {
		(*list)[j].Add(copyData[j-1], copyData[j])
	}
	*list = append(*list, big.NewInt(1))
}

func printSpace(writer *bufio.Writer, prefixSLen int, times int) {
	buffer := bytes.NewBufferString("")
	for i := 0; i < prefixSLen*times; i++ {
		buffer.WriteString(" ")
	}
	_, _ = writer.WriteString(buffer.String())
	_ = writer.Flush()
}

func computeSpaceLen(numLen int) int {
	if numLen < 4 {
		return 3
	}
	if numLen%2 == 0 {
		return numLen + 1
	}
	return numLen
}

func computeMaxLen(level int) int {
	var m int
	if level%2 == 1 {
		m = level / 2
	} else {
		m = level/2 - 1
	}
	diBy, _ := factorial(level - 1)
	f1, _ := factorial(level / 2)
	f2, _ := factorial(m)
	diBy.Div(diBy, f1.Mul(f1, f2))
	return len(diBy.String())
}

func factorial(n int) (result *big.Int, err error) {
	if n < 0 {
		err = errors.New("参数不合法：" + strconv.Itoa(n))
	}
	result = big.NewInt(int64(n))
	for i := n - 1; i > 1; i-- {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result, nil
}

func getStdInNum() int {
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		return 0
	}
	i, err := strconv.Atoi(string(line))
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return i
}
