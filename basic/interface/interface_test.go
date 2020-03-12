package _interface

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {

	s := "sad{{}{}{}"
	convert(s)
}

func convert(s interface{}) {

	fmt.Println(s.(string))
}
