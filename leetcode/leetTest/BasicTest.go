package main

import (
	"fmt"
	leet "goSecond/leetcode"
)

func main() {
	q6()
}

func q3() {
	result := map[string]int{"aab": 2, "abcdaec": 5, "abcd": 4, "aabb": 2, "啊啊表": 2, "一二三四一五三": 5}
	var b bool
	for k, v := range result {
		if b = leet.LengthOfLongestSubstring(k) != v; b {
			fmt.Println(k, ":", v)
		}
	}
	if !b {
		fmt.Println("ok")
	}
}

func q4() {
	var num1 [3]int = [...]int{1, 2, 3}
	var num2 = [...]int{1, 4, 5, 7, 9}

	i := leet.FindMedianSortedArrays(num1[:], num2[:])
	fmt.Println(i)
}

func q5() {
	s := "abcddce"
	s2 := leet.LongestPalindrome(s)
	fmt.Println(s2)
}

func q6() {
	s := "LEETCODEISHIRING"
	result := map[int]string{3: "LCIRETOESIIGEDHN", 4: "LDREOEIIECIHNTSG"}
	var b bool
	for k, v := range result {
		convert := leet.Convert(s, k)
		if b = convert != v; b {
			fmt.Println(k, ":", v, "my answer :", convert)
		}
	}
	if !b {
		fmt.Println("ok")
	}
}

func q6a() {
	s := "ABCDE"
	convert := leet.Convert(s, 4)
	fmt.Println(convert)
}
