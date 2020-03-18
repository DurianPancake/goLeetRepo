package model

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
