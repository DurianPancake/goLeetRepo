package utils

import (
	"fmt"
	"goSecond/hello/fun/formula/model"
	"strings"
)

const (
	base    = "base"
	logic   = "logic"
	bracket = "bracket"
)

//
func GenerateCondition(from string) (json string, err error) {

	//TODO 括号校验

	// 预处理
	from = strings.ReplaceAll(from, " ", "")
	fmt.Println(from)

	// 生成对象
	logicalUnit, err := generateLogicalGroupFromText(from)

	//TODO 生成目标表达式字符串
	if err != nil {
		return "", err
	}
	return generateJson(logicalUnit), nil
}

// 生成逻辑表达式对象
// 先递归的解析为栈组
// 在同一层级的栈组中根据存在的优先级生成对象
// 再递归的返还生成的逻辑组对象到最高层
func generateLogicalGroupFromText(from string) (model.LogicalGroup, error) {

	// 解析为栈组
	stackGroup := getStackGroup(from)
	fmt.Println(stackGroup.Stacks)
	defer func() {

	}()
	// 生成对象
	newLogicalGroupByStacks(stackGroup)
	return model.LogicalGroup{}, nil
}

// 生成栈组
func getStackGroup(from string) *model.StackGroup {

	currentGroup := new(model.StackGroup)
	stacks := make([]model.LogicalStack, 0)
	runes := []rune(from)
	cursor, valueStartCursor := 0, 0
	for cursor < len(runes) {
		operate := model.MatchOperate(from, cursor)
		if operate == model.Null {
			cursor++
			continue
		}
		switch operate.Type {
		case base:
			// 找到基本运算符
			if cursor == valueStartCursor {
				panic("field not found")
			}
			field := string(runes[valueStartCursor:cursor])
			valueIndex := cursor + len(operate.Symbol)
			var nextCursor int
			for nextCursor = valueIndex; nextCursor < len(runes); nextCursor++ {
				op := model.MatchOperate(from, nextCursor)
				if op != model.Null {
					break
				}
			}
			value := string(runes[valueIndex:nextCursor])
			stacks = append(stacks, model.LogicalStack{
				Kind:     1,
				Operator: model.Operator{},
				Unit: model.Unit{
					Field:   field,
					Operate: operate,
					Value:   value,
				},
				StackGroup: model.StackGroup{},
			})
			cursor, valueStartCursor = nextCursor, nextCursor

		case logic:
			// 找到逻辑运算符
			cursor += len(operate.Symbol)
			stacks = append(stacks, model.LogicalStack{
				Kind:       0,
				Operator:   operate,
				Unit:       model.Unit{},
				StackGroup: model.StackGroup{},
			})
			valueStartCursor = cursor

		case bracket:
			// 找到括号
			nextCursor := lastRuneIndex(from, ')') + 1
			subStr := string(runes[cursor+1 : nextCursor-1])
			fmt.Println(subStr)
			group := getStackGroup(subStr)
			stacks = append(stacks, model.LogicalStack{
				Kind:       2,
				Operator:   model.Operator{},
				Unit:       model.Unit{},
				StackGroup: *group,
			})
			cursor, valueStartCursor = nextCursor, nextCursor
		}
	}
	currentGroup.Stacks = stacks
	return currentGroup
}

// 生成JSON表达式
func generateJson(logicalUnit model.LogicalGroup) string {

	return ""
}

// 中文字符索引
func lastRuneIndex(text string, char rune) int {
	runes := []rune(text)
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == char {
			return i
		}
	}
	return -1
}

// 生成对象
// 优先级：组 > ! > && > ||
func newLogicalGroupByStacks(group *model.StackGroup) *model.LogicalGroup {

	// 初始化
	current := new(model.LogicalGroup)
	subGroups := make([]model.LogicalGroup, 0)
	_ = make([]model.Unit, 0)

	stacks := group.Stacks
	// 序号到优先级的映射字典
	sortMap := make(map[int]int, 0)
	// 遍历时排序
	for i, stack := range stacks {
		sortMap[i] = stack.Kind
	}

	for i, stack := range stacks {
		if stack.Kind == 2 {
			delete(sortMap, i)
			logicalGroup := newLogicalGroupByStacks(&stack.StackGroup)
			subGroups = append(subGroups, *logicalGroup)
		}
	}

	return current
}
