package fomular

import (
	"fmt"
	"strings"
	"testing"
)

var (
	f1 = "A=3 && B!=4 ||  C>5"
	f2 = "A=3 || B!=4 && C>5"
)

func TestStack(t *testing.T) {

	f1 = strings.ReplaceAll(f1, " ", "")
	fmt.Println(f1) //A=3&&B!=4||C>5

	stacks := make([]Stack, 0)

	runes := []rune(f1)
	lastIndex := 0
	for cursor := 0; cursor < len(runes); cursor++ {
		offsets := MatchOperator(runes[cursor])
		// 符号
		if len(offsets) > 0 {
			maxOffset, index := 0, -1
			var op Operator
			for i, offset := range offsets {
				// 最大匹配
				if operator := Match(string(runes[cursor : cursor+offset+1])); operator != "" {
					if offset >= maxOffset {
						index = i
						maxOffset = offset
						op = operator
					}
				}
			}
			if index < 0 {
				continue
			}
			// 赋值
			if lastIndex != cursor {
				stacks = append(stacks, Stack{
					Kind:    Field,
					Field:   string(runes[lastIndex:cursor]),
					Operate: "",
					Value:   "",
				})
			}
			stacks = append(stacks, Stack{
				Kind:    Operate,
				Field:   "",
				Operate: op,
				Value:   "",
			})
			lastIndex = cursor
			cursor += maxOffset
		}
	}

	fmt.Println(stacks)
}
