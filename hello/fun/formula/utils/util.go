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
			nextCursor := model.FindMatchBracketIndex(from, cursor) + 1
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

// 生成对象
// 优先级：组 > ! > && > ||
func newLogicalGroupByStacks(group *model.StackGroup) *model.LogicalGroup {

	// 初始化
	current := new(model.LogicalGroup)
	subUnits := make([]model.Unit, 0)

	stacks := group.Stacks
	// 最简情况
	if len(stacks) == 1 {
		if stacks[0].Kind == 1 {
			subUnits = append(subUnits, stacks[0].Unit)
			current.Units = subUnits
			return current
		} else if stacks[0].Kind == 2 {
			return newLogicalGroupByStacks(&stacks[0].StackGroup)
		}
		panic("expression error")
	}

	mergeMap := make(map[int]model.LogicalStack, 0)
	for i, stack := range stacks {
		mergeMap[i] = stack
	}

	subGroups := make([]model.LogicalGroup, 0)
	// 获取运算符
	symbolList := group.GetPreferenceList()
	symbol, pos, all := getSameLevelSymbols(symbolList)
	// TODO continue

	for len(symbolList) > 1 {
		symbol := symbolList[0]
		stack := symbol.Stack
		symIndex := symbol.Index
		switch stack.Symbol {
		case "!":
			nextIndex := symIndex + 1
			if len(stacks) <= nextIndex {
				panic(fmt.Sprintf("wrong operator: %s", stack.Symbol))
			}
			logicalStack := stacks[nextIndex]
			if logicalStack.Kind == 0 {
				panic(fmt.Sprintf("wrong operator: %s", stack.Symbol))
			}
		}
	}

	return current
}

func getSameLevelSymbols(list []model.Preference) (symbol string, nextPos int, all bool) {
	stack := list[0].Stack
	symbol = stack.Symbol
	for i, preference := range list {
		if preference.Stack.Symbol != symbol {
			return symbol, i, false
		}
	}
	return symbol, len(list), true
}
