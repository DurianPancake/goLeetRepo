package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

//
func GenerateCondition(from string) (json string, err error) {

	//1.校验
	err = validateText(from)
	if err != nil {
		return "", err
	}
	// 2.预处理
	from = strings.ReplaceAll(from, " ", "")
	fmt.Println(from)

	// 3.生成对象
	logicalUnit, err := generateLogicalGroupFromText(from)
	if err != nil {
		return "", err
	}
	// 4.生成JSON字符串
	return generateJson(logicalUnit)
}

// TODO 校验 来源字符串
// 校验1. 括号数量
// 校验2. 括号顺序
// 校验3. 括号中有值
func validateText(from string) error {
	return nil
}

// 生成逻辑表达式对象
// 先递归的解析为栈组
// 在同一层级的栈组中根据存在的优先级生成对象
// 再递归的返还生成的逻辑组对象到最高层
func generateLogicalGroupFromText(from string) (LogicalGroup, error) {

	// 解析为栈、】0-----0
	stackGroup := getStackGroup(from)
	fmt.Println(stackGroup.Stacks)

	// 生成对象
	logicalGroup := newLogicalGroupByStacks(stackGroup)
	return *logicalGroup, nil
}

// 生成栈组
func getStackGroup(from string) *StackGroup {

	currentGroup := new(StackGroup)
	stacks := make([]LogicalStack, 0)
	runes := []rune(from)
	cursor, valueStartCursor := 0, 0
	for cursor < len(runes) {
		operate := matchOperate(from, cursor)
		if operate == Null {
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
				op := matchOperate(from, nextCursor)
				if op != Null {
					break
				}
			}
			value := string(runes[valueIndex:nextCursor])
			stacks = append(stacks, LogicalStack{
				Kind:     1,
				Operator: Operator{},
				Unit: Unit{
					Field:   field,
					Operate: operate,
					Value:   value,
				},
				StackGroup: StackGroup{},
			})
			cursor, valueStartCursor = nextCursor, nextCursor

		case logic:
			// 找到逻辑运算符
			cursor += len(operate.Symbol)
			stacks = append(stacks, LogicalStack{
				Kind:       0,
				Operator:   operate,
				Unit:       Unit{},
				StackGroup: StackGroup{},
			})
			valueStartCursor = cursor

		case bracket:
			// 找到括号
			nextCursor := findMatchBracketIndex(from, cursor) + 1
			subStr := string(runes[cursor+1 : nextCursor-1])
			fmt.Println(subStr)
			group := getStackGroup(subStr)
			stacks = append(stacks, LogicalStack{
				Kind:       2,
				Operator:   Operator{},
				Unit:       Unit{},
				StackGroup: *group,
			})
			cursor, valueStartCursor = nextCursor, nextCursor
		}
	}
	currentGroup.Stacks = stacks
	return currentGroup
}

// 生成JSON表达式
// TODO 暂定生成
func generateJson(logicalUnit LogicalGroup) (string, error) {
	marshal, err := json.Marshal(logicalUnit)
	return string(marshal), err
}

// 生成对象
// 优先级：组 > ! > && > ||
// 会抛出panic
func newLogicalGroupByStacks(group *StackGroup) *LogicalGroup {

	stacks := group.Stacks
	// 最简情况
	if len(stacks) == 1 {
		if stacks[0].Kind == 1 {
			// 初始化
			current := newLogicGroupByStack(stacks[0])
			return current
		} else if stacks[0].Kind == 2 {
			return newLogicalGroupByStacks(&stacks[0].StackGroup)
		}
		panic("expression error")
	}
	// 封装一层栈组，因为栈的入栈顺序有意义
	prefList := getPreListFromStacks(stacks)
	// 获取运算符
	symbolList := group.GetPreferenceList()
	// 根据栈组，排好序的符号列表合并栈到一条以生成逻辑运算对象
	return mergeStack(prefList, symbolList)
}

// 从Stack创建LogicalGroup，只允许类型为Unit的创建
func newLogicGroupByStack(stack LogicalStack) *LogicalGroup {
	if stack.Kind != 1 {
		panic("error Logical Group init from stack: stack type is not right")
	}
	current := new(LogicalGroup)
	subUnits := make([]Unit, 0)
	subUnits = append(subUnits, stack.Unit)
	current.Units = subUnits
	return current
}

// 合并栈组（StackGroup）为逻辑计算组（LogicalGroup）
// 根据运算符列表，找出同一类型的运算符，它们具有相同的优先级
// 如果这些运算符在计算位置上连续，连续的部分可以合成一个组，不连续的部分自成一组
// 运算符运算会消耗基本单元（Unit）和组（StackGroup）和它自身（Symbol）
// 因此链表会因为计算不断产生新组（LogicalGroup），统一用Preference封装起来
// 同时栈组（StackGroup）也会不断减少
// 如果有下一优先级的运算符列表，重复以上过程
// 当运算符消耗完毕时，计算式只会留下一组（LogicalGroup）将其返回
func mergeStack(prefList *List, symbolList []Preference) *LogicalGroup {

	// 同类型计算符
	for prefList.Length() != 1 {
		symbol, pos, all := getSameLevelSymbols(symbolList)
		samePos := symbolList[0:pos]
		if !all {
			symbolList = symbolList[pos:]
		}

		prefMap := make(map[int]Preference, 0)
		for _, po := range samePos {
			prefMap[po.Index] = po
		}
		// 分组
		// map[分组][序号集]
		seriesGroup := make(map[int][]int, 0)
		lastIndex := 0
		for _, symbol := range samePos {
			orders := seriesGroup[lastIndex]
			if orders == nil {
				orders = make([]int, 0)
				seriesGroup[lastIndex] = orders
			}
			bigOrder := len(orders) - 1
			if bigOrder < 0 {
				orders = append(orders, symbol.Index)
				seriesGroup[lastIndex] = orders
				continue
			}
			isFarAway := orders[bigOrder]-symbol.Index > 2
			if !isFarAway {
				orders = append(orders, symbol.Index)
				seriesGroup[lastIndex] = orders
			} else if isFarAway {
				lastIndex++
				nextGroup := make([]int, 0)
				nextGroup = append(nextGroup, symbol.Index)
				seriesGroup[lastIndex] = nextGroup
			}
		}
		// 根据运算规则合并产生新组，一个组内为同一类型
		for _, orders := range seriesGroup {
			//
			subUnits := make([]Unit, 0)
			subGroups := make([]LogicalGroup, 0)
			logicUnit := LogicalGroup{
				Operator: Operator{Symbol: symbol},
			}

			switch symbol {
			case "!":
				// 非运算只取order后一位
				// 一个组的非运算涉及到的stack个数为 len(order) * 2个
				minIndex := prefList.Length()
				for _, order := range orders {
					// 拿到下一个栈帧
					cur := prefList.Head()
					index := 0
					for cur.Data.(Preference).Index != order {
						cur = cur.Next
						index++
					}
					if index < minIndex {
						minIndex = index
					}
				}
				// 添加方法
				addSub(prefList, &subGroups, &subUnits, minIndex+1, len(orders))
				logicUnit.Units = subUnits
				logicUnit.LogicalUnits = subGroups
				// 删除和添加
				for i := 0; i < len(orders)*2; i++ {
					prefList.RemoveAtIndex(minIndex)
				}
				prefList.Insert(Preference{
					Index: minIndex,
					Kind:  1,
					Group: logicUnit,
				}, minIndex)
			case "&&":
				fallthrough
			case "||":
				// 与运算和或运算需要取前一位和后一位
				// 一个组的非运算涉及到的stack个数为 len(order)*2 + 1个
				minIndex := prefList.Length()
				for _, order := range orders {
					// 拿到下一个栈帧
					cur := prefList.Head()
					index := 0
					for cur.Data.(Preference).Index != order {
						cur = cur.Next
						index++
					}
					if index < minIndex {
						minIndex = index
					}
				}
				// 添加方法
				alternativePos := minIndex - 1
				addSub(prefList, &subGroups, &subUnits, alternativePos, len(orders)+1)
				logicUnit.Units = subUnits
				logicUnit.LogicalUnits = subGroups
				// 删除和添加
				for i := 0; i < len(orders)*2+1; i++ {
					prefList.RemoveAtIndex(alternativePos)
				}
				prefList.Insert(Preference{
					Index: alternativePos,
					Kind:  1,
					Group: logicUnit,
				}, alternativePos)
			default:
			}
		}
	}
	// 这里有可能是Group，但也有可能是只有一层括号的Unit
	finalPreference := prefList.headNode.Data.(Preference)
	var final LogicalGroup
	if finalPreference.Group.equals(&LogicalGroup{}) {
		final = *newLogicGroupByStack(finalPreference.Stack)
	} else {
		final = finalPreference.Group
	}
	return &final
}

// 添加元素
func addSub(prefList *List, groups *[]LogicalGroup, units *[]Unit, index int, times int) {
	cur := prefList.get(index)
	for i := 0; i < times; i++ {
		stackPref := cur.Data.(Preference)
		if stackPref.Kind == 0 {
			if stackPref.Stack.Kind == 0 {
				panic(fmt.Sprintf("expression error"))
			} else if stackPref.Stack.Kind == 2 {
				group := stackPref.Stack.StackGroup
				subPreference := group.GetPreferenceList()
				// 内层退出条件
				logicalGroup := *mergeStack(getPreListFromStacks(group.Stacks), subPreference)
				if logicalGroup.Symbol == "" {
					*units = append(*units, logicalGroup.Units[0])
				} else {
					*groups = append(*groups, logicalGroup)
				}
			} else {
				*units = append(*units, stackPref.Stack.Unit)
			}
		} else if stackPref.Kind == 1 {
			*groups = append(*groups, stackPref.Group)
		}
		if i < times-1 {
			cur = cur.Next.Next
		}
	}
}

// 获取同优先级符号集合
func getSameLevelSymbols(list []Preference) (symbol string, nextPos int, all bool) {
	stack := list[0].Stack
	symbol = stack.Symbol
	for i, preference := range list {
		if preference.Stack.Symbol != symbol {
			return symbol, i, false
		}
	}
	return symbol, len(list), true
}

func getPreListFromStacks(stacks []LogicalStack) *List {
	prefList := new(List)
	for i, stack := range stacks {
		prefList.Append(Preference{
			Index: i,
			Kind:  0,
			Stack: stack,
		})
	}
	return prefList
}
