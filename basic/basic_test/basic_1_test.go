package basic_test

import (
	"fmt"
	"goSecond/basic/constant"
	"testing"
)

func TestConst(t *testing.T) {
	fmt.Println(constant.Monday,
		constant.Tuesday,
		constant.Wednesday,
		constant.Thursday,
		constant.Friday,
		constant.Saturday,
		constant.Sunday)
}

func TestMonth(t *testing.T) {
	fmt.Println(constant.January, constant.February)
}
