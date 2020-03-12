package pkgs

import (
	"fmt"
	"regexp"
	"testing"
)

var strGet = "$134{get}"

func TestGet(t *testing.T) {

	exp := regexp.MustCompile(`\$(.+){(.+)}`)

	submatch := exp.FindAllStringSubmatch(strGet, -1)
	fmt.Println(submatch)

	for i, strings := range submatch {
		fmt.Println(i, strings)
		fmt.Println("group 0", strings[0])
		fmt.Println("group 1", strings[1])
		fmt.Println("group 2", strings[2])
	}
}
