package main

import (
	"fmt"
	leet "goSecond/leetcode/q1_20"
	"testing"
)

func TestSet2(t *testing.T) {
	q11()
}

func q11() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := leet.Q11MaxArea(height)
	fmt.Println(area)
}
