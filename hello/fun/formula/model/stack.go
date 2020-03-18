package model

import "goSecond/hello/fun/formula/utils"

// 解析时运算栈
type LogicalStack struct {
	Kind     int // 0\1\2 依次表示Operator、Unit、LogicalGroup
	Operator     // 只会是逻辑运算符
	Unit
	StackGroup
}

// 解析时括号类型运算组
type StackGroup struct {
	Stacks []LogicalStack
}

type Preference struct {
	Index int
	Stack LogicalStack
}

// 获取优先级排序列表
// 只纪录有逻辑运算符
func (sg *StackGroup) GetPreferenceList() []Preference {

	list := new(utils.List)
	stacks := sg.Stacks
	for i, stack := range stacks {
		if stack.Kind != 0 {
			continue
		}
		index := 0
		for node := list.Head(); node.HasNext(); index++ {
			preference := node.Data.(Preference)
			if preference.Stack.compareTo(&stack) > 0 {
				node = node.Next
				continue
			} else {
				break
			}
		}
		list.Insert(Preference{
			Index: i,
			Stack: stack,
		}, index)
	}
	ps := make([]Preference, list.Length())
	index := 0
	for node := list.Head(); node.HasNext(); index++ {
		ps[index] = node.Data.(Preference)
	}
	return ps
}

var symbols = map[string]int{
	"!":  2,
	"&&": 1,
	"||": 0,
}

// 比较方法: 组 > ! > && > ||
// a.compareTo(b) if a > b return positive number
func (ls *LogicalStack) compareTo(stack *LogicalStack) int {
	if stack.Kind == 2 {
		return ls.Kind - 2
	}
	if stack.Kind == 0 {
		if ls.Kind == 2 {
			return 1
		}
		if ls.Kind == 1 {
			return -1
		}
		// compare symbol
		return symbols[ls.Symbol] - symbols[stack.Symbol]
	}
	if stack.Kind == 1 {
		if ls.Kind == 1 {
			return 0
		}
		return 1
	}
	return 0
}
