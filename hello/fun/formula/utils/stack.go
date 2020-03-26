package utils

// 解析时运算栈
type logicalStack struct {
	kind     int // 0\1\2 依次表示Operator、unit、logicalGroup
	operator     // 只会是逻辑运算符
	unit
	stackGroup
}

// 解析时括号类型运算组
type stackGroup struct {
	stacks []logicalStack
}

type preference struct {
	index int
	kind  int // 0\1 表示Stack group
	stack logicalStack
	group logicalGroup
}

// 获取优先级排序列表
// 只纪录有逻辑运算符
func (sg *stackGroup) getPreferenceList() []preference {

	list := new(List)
	stacks := sg.stacks
	for i, stack := range stacks {
		if stack.kind != 0 {
			continue
		}
		index := 0
		for node := list.Head(); node != nil; index++ {
			preference := node.Data.(preference)
			if preference.stack.compareTo(&stack) > 0 {
				node = node.Next
				continue
			} else {
				break
			}
		}
		list.Insert(preference{
			index: i,
			stack: stack,
		}, index)
	}
	ps := make([]preference, list.Length())
	index := 0
	for node := list.Head(); node != nil; index++ {
		ps[index] = node.Data.(preference)
		node = node.Next
	}
	return ps
}

var symbols = map[string]int{
	not.symbol: 2,
	and.symbol: 1,
	or.symbol:  0,
}

// 比较方法: 组 > ! > && > ||
// a.compareTo(b) if a > b return positive number
func (ls *logicalStack) compareTo(stack *logicalStack) int {
	if stack.kind == 2 {
		return ls.kind - 2
	}
	if stack.kind == 0 {
		if ls.kind == 2 {
			return 1
		}
		if ls.kind == 1 {
			return -1
		}
		// compare symbol
		return symbols[ls.symbol] - symbols[stack.symbol]
	}
	if stack.kind == 1 {
		if ls.kind == 1 {
			return 0
		}
		return 1
	}
	return 0
}
