package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

//
func GenerateCondition(from string) (str string, err error) {

	//1.校验
	err = validateText(from)
	if err != nil {
		return "", err
	}
	// 2.预处理
	from = strings.ReplaceAll(from, " ", "")

	// 3.生成对象
	logicalUnit, err := generateLogicalGroupFromText(from)
	if err != nil {
		return "", err
	}
	// 4.生成JSON字符串
	model := generateExportModel(logicalUnit)
	bytes, err := json.Marshal(model)
	if err != nil {
		return "{}", errors.New("生成Json异常")
	}
	return string(bytes), err
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
func generateLogicalGroupFromText(from string) (logicalGroup, error) {

	// 解析为栈、】0-----0
	stackGroup := getStackGroup(from)
	//fmt.Println(stackGroup.stacks)

	// 生成对象
	logicalGroup := newLogicalGroupByStacks(stackGroup)
	return *logicalGroup, nil
}

// 生成栈组
func getStackGroup(from string) *stackGroup {

	currentGroup := new(stackGroup)
	stacks := make([]logicalStack, 0)
	runes := []rune(from)
	cursor, valueStartCursor := 0, 0
	for cursor < len(runes) {
		operate := matchOperate(from, cursor)
		if operate == null {
			cursor++
			continue
		}
		switch operate._type {
		case base:
			// 找到基本运算符
			if cursor == valueStartCursor {
				panic("field not found")
			}
			field := string(runes[valueStartCursor:cursor])
			valueIndex := cursor + len(operate.symbol)
			var nextCursor int
			for nextCursor = valueIndex; nextCursor < len(runes); nextCursor++ {
				op := matchOperate(from, nextCursor)
				if op != null {
					break
				}
			}
			value := string(runes[valueIndex:nextCursor])
			stacks = append(stacks, logicalStack{
				kind:     1,
				operator: operator{},
				unit: unit{
					field:   field,
					operate: operate,
					value:   value,
				},
				stackGroup: stackGroup{},
			})
			cursor, valueStartCursor = nextCursor, nextCursor

		case logic:
			// 找到逻辑运算符
			cursor += len(operate.symbol)
			stacks = append(stacks, logicalStack{
				kind:       0,
				operator:   operate,
				unit:       unit{},
				stackGroup: stackGroup{},
			})
			valueStartCursor = cursor

		case bracket:
			// 找到括号
			nextCursor := findMatchBracketIndex(from, cursor) + 1
			subStr := string(runes[cursor+1 : nextCursor-1])
			//fmt.Println(subStr)
			group := getStackGroup(subStr)
			stacks = append(stacks, logicalStack{
				kind:       2,
				operator:   operator{},
				unit:       unit{},
				stackGroup: *group,
			})
			cursor, valueStartCursor = nextCursor, nextCursor
		}
	}
	currentGroup.stacks = stacks
	return currentGroup
}

// 生成对象
// 优先级：组 > ! > && > ||
// 会抛出panic
func newLogicalGroupByStacks(group *stackGroup) *logicalGroup {

	stacks := group.stacks
	// 最简情况
	if len(stacks) == 1 {
		if stacks[0].kind == 1 {
			// 初始化
			current := newLogicGroupByStack(stacks[0])
			return current
		} else if stacks[0].kind == 2 {
			return newLogicalGroupByStacks(&stacks[0].stackGroup)
		}
		panic("expression error")
	}
	// 封装一层栈组，因为栈的入栈顺序有意义
	prefList := getPreListFromStacks(stacks)
	// 获取运算符
	symbolList := group.getPreferenceList()
	// 根据栈组，排好序的符号列表合并栈到一条以生成逻辑运算对象
	return mergeStack(prefList, symbolList)
}

// 从Stack创建LogicalGroup，只允许类型为Unit的创建
func newLogicGroupByStack(stack logicalStack) *logicalGroup {
	if stack.kind != 1 {
		panic("error Logical group init from stack: stack type is not right")
	}
	current := new(logicalGroup)
	subUnits := make([]unit, 0)
	subUnits = append(subUnits, stack.unit)
	current.units = subUnits
	return current
}

// 合并栈组（stackGroup）为逻辑计算组（logicalGroup）
// 根据运算符列表，找出同一类型的运算符，它们具有相同的优先级
// 如果这些运算符在计算位置上连续，连续的部分可以合成一个组，不连续的部分自成一组
// 运算符运算会消耗基本单元（unit）和组（stackGroup）和它自身（symbol）
// 因此链表会因为计算不断产生新组（logicalGroup），统一用Preference封装起来
// 同时栈组（stackGroup）也会不断减少
// 如果有下一优先级的运算符列表，重复以上过程
// 当运算符消耗完毕时，计算式只会留下一组（logicalGroup）将其返回
func mergeStack(prefList *List, symbolList []preference) *logicalGroup {

	// 同类型计算符
	for prefList.Length() != 1 {
		symbol, pos, all := getSameLevelSymbols(symbolList)
		samePos := symbolList[0:pos]
		if !all {
			symbolList = symbolList[pos:]
		}

		prefMap := make(map[int]preference, 0)
		for _, po := range samePos {
			prefMap[po.index] = po
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
				orders = append(orders, symbol.index)
				seriesGroup[lastIndex] = orders
				continue
			}
			isFarAway := orders[bigOrder]-symbol.index > 2
			if !isFarAway {
				orders = append(orders, symbol.index)
				seriesGroup[lastIndex] = orders
			} else if isFarAway {
				lastIndex++
				nextGroup := make([]int, 0)
				nextGroup = append(nextGroup, symbol.index)
				seriesGroup[lastIndex] = nextGroup
			}
		}
		// 根据运算规则合并产生新组，一个组内为同一类型
		for _, orders := range seriesGroup {
			//
			subUnits := make([]unit, 0)
			subGroups := make([]logicalGroup, 0)
			logicUnit := logicalGroup{
				operator: operator{symbol: symbol},
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
					for cur.Data.(preference).index != order {
						cur = cur.Next
						index++
					}
					if index < minIndex {
						minIndex = index
					}
				}
				// 添加方法
				addSub(prefList, &subGroups, &subUnits, minIndex+1, len(orders))
				logicUnit.units = subUnits
				logicUnit.logicalUnits = subGroups
				// 删除和添加
				for i := 0; i < len(orders)*2; i++ {
					prefList.RemoveAtIndex(minIndex)
				}
				prefList.Insert(preference{
					index: minIndex,
					kind:  1,
					group: logicUnit,
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
					for cur.Data.(preference).index != order {
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
				logicUnit.units = subUnits
				logicUnit.logicalUnits = subGroups
				// 删除和添加
				for i := 0; i < len(orders)*2+1; i++ {
					prefList.RemoveAtIndex(alternativePos)
				}
				prefList.Insert(preference{
					index: alternativePos,
					kind:  1,
					group: logicUnit,
				}, alternativePos)
			default:
			}
		}
	}
	// 这里有可能是Group，但也有可能是只有一层括号的Unit
	finalPreference := prefList.headNode.Data.(preference)
	var final logicalGroup
	if finalPreference.group.equals(&logicalGroup{}) {
		final = *newLogicGroupByStack(finalPreference.stack)
	} else {
		final = finalPreference.group
	}
	return &final
}

// 添加元素
func addSub(prefList *List, groups *[]logicalGroup, units *[]unit, index int, times int) {
	cur := prefList.get(index)
	for i := 0; i < times; i++ {
		stackPref := cur.Data.(preference)
		if stackPref.kind == 0 {
			if stackPref.stack.kind == 0 {
				panic(fmt.Sprintf("expression error"))
			} else if stackPref.stack.kind == 2 {
				group := stackPref.stack.stackGroup
				subPreference := group.getPreferenceList()
				// 内层退出条件
				logicalGroup := *mergeStack(getPreListFromStacks(group.stacks), subPreference)
				if logicalGroup.symbol == "" {
					*units = append(*units, logicalGroup.units[0])
				} else {
					*groups = append(*groups, logicalGroup)
				}
			} else {
				*units = append(*units, stackPref.stack.unit)
			}
		} else if stackPref.kind == 1 {
			*groups = append(*groups, stackPref.group)
		}
		if i < times-1 {
			cur = cur.Next.Next
		}
	}
}

// 获取同优先级符号集合
func getSameLevelSymbols(list []preference) (symbol string, nextPos int, all bool) {
	stack := list[0].stack
	symbol = stack.symbol
	for i, preference := range list {
		if preference.stack.symbol != symbol {
			return symbol, i, false
		}
	}
	return symbol, len(list), true
}

func getPreListFromStacks(stacks []logicalStack) *List {
	prefList := new(List)
	for i, stack := range stacks {
		prefList.Append(preference{
			index: i,
			kind:  0,
			stack: stack,
		})
	}
	return prefList
}

// 生成JSON表达式
func generateExportModel(logicalUnit logicalGroup) interface{} {

	// 最简情况
	if logicalUnit.operator.symbol == "" {
		u := logicalUnit.units[0]
		return generateJsonFromUnit(u)
	}
	conditions := make([]interface{}, 0)
	// units的部分
	units := logicalUnit.units
	for _, u := range units {
		jsonFromUnit := generateJsonFromUnit(u)
		conditions = append(conditions, jsonFromUnit)
	}
	// logicalGroup的部分
	if logicalUnits := logicalUnit.logicalUnits; logicalUnits != nil && len(logicalUnits) > 0 {
		for _, group := range logicalUnits {
			groupJson := generateExportModel(group)
			conditions = append(conditions, groupJson)
		}
	}

	group := UnitGroup{
		Type:       logs[logicalUnit.operator.symbol],
		Conditions: conditions,
	}
	return group
}

func generateJsonFromUnit(u unit) Unit {
	return Unit{
		Field:    u.field,
		Operator: u.operate.symbol,
		Value:    u.value,
		Type:     condition,
	}
}
