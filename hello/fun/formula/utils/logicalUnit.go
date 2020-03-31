package utils

const (
	base    = "base"
	logic   = "logic"
	bracket = "bracket"
)

// 所有符号的定义
// symbol：代表着可以在公式中识别的符号字符串。
// 可以在此处替换字符串而不会影响程序运行，symbol不建议使用空格，因为与预处理方式冲突
// 但可能会存在的影响例如将'&&'改为'&'，公式中如果存在&&的运算符将会识别为连续两个'&'与运算符
// type：代表符号类型
// TODO:将符号字符串提供可配置化
var (
	eq    = operator{"=", base}
	ne    = operator{"!=", base}
	gt    = operator{">", base}
	lt    = operator{"<", base}
	gte   = operator{">=", base}
	lte   = operator{"<=", base}
	like  = operator{"like", base}
	in    = operator{"in", base}
	notIn = operator{"not_in", base}
	//
	not = operator{"!", logic}
	and = operator{"&&", logic}
	or  = operator{"||", logic}
	//
	leftBracket  = operator{"(", bracket}
	rightBracket = operator{")", bracket}
	//
	null = operator{"", ""}
)

// 运算符
type operator struct {
	symbol string
	_type  string
}

// 基本逻辑运算单元，代表一个布尔值类型
type unit struct {
	field   string
	operate operator
	value   interface{}
}

// 描述一组逻辑运算关系，在该组中，逻辑运算符是一样的，表示同一优先级的运算
// 基本单位可以是Unit或是Logical
// ! > && > ||
// !运算符只能允许存在一个Unit或LogicalGroup
type logicalGroup struct {
	operator
	units        []unit
	logicalUnits []logicalGroup
}

var operators = []operator{
	eq,
	ne,
	gt,
	lt,
	gte,
	lte,
	like,
	in,
	notIn,
	and,
	or,
	not,
	leftBracket,
	rightBracket,
	null,
}

// 最大匹配原则，尽可能长的匹配字段或符号
// 获取每个匹配的偏移量，选取偏移量最大的
func matchOperate(text string, index int) operator {

	runes := []rune(text)
	maxOffset := 0
	var chosenOp operator
	for _, operator := range operators {
		if operator == null {
			continue
		}
		firstChar := []rune(operator.symbol)[0]
		if firstChar == runes[index] {
			if length := len(operator.symbol); length > maxOffset && string(runes[index:index+length]) == operator.symbol {
				maxOffset = length
				chosenOp = operator
			}
		}
	}
	return chosenOp
}

// 寻找匹配右括号的索引位置
// ((() ())())
func findMatchBracketIndex(text string, index int) int {

	operate := matchOperate(text, index)
	if operate._type != bracket || operate.symbol != leftBracket.symbol {
		return -1
	}
	//
	count := 1
	runes := []rune(text)
	leftSymbolChs := []rune(leftBracket.symbol)
	rightSymbolChs := []rune(rightBracket.symbol)
	for i := index + 1; i < len(runes); i++ {
		var leftJudge, rightJudge = true, true
		for m := 0; m < len(leftSymbolChs); m++ {
			if i+m == len(runes) || runes[i+m] != leftSymbolChs[m] {
				leftJudge = false
				break
			}
		}
		if leftJudge {
			count += 1
		}
		for m := 0; m < len(rightSymbolChs); m++ {
			if i+m == len(runes) || runes[i+m] != rightSymbolChs[m] {
				rightJudge = false
				break
			}
		}
		if rightJudge {
			count -= 1
			if count == 0 {
				return i
			}
		}
	}
	return -1
}

func (op1 *operator) equals(op2 *operator) bool {
	return op1.symbol == op2.symbol
}

func (u1 *unit) equals(u2 *unit) bool {
	return u1.operate.equals(&u2.operate) && u1.field == u2.field && u1.value == u2.value
}

// 判断逻辑组是否相等
func (lg1 *logicalGroup) equals(lg2 *logicalGroup) bool {
	return lg1.symbol == lg2.symbol && isUnitsEquals(lg1, lg2) && isLogicalUnitsEquals(lg1, lg2)
}

func isLogicalUnitsEquals(lg1 *logicalGroup, lg2 *logicalGroup) bool {
	if len(lg1.logicalUnits) != len(lg2.logicalUnits) {
		return false
	}
	for i, unit := range lg1.logicalUnits {
		units := lg2.logicalUnits
		if !units[i].equals(&unit) {
			return false
		}
	}
	return true
}

func isUnitsEquals(lg1 *logicalGroup, lg2 *logicalGroup) bool {
	if len(lg1.units) != len(lg2.units) {
		return false
	}
	for i, unit := range lg1.units {
		units := lg2.units
		if !units[i].equals(&unit) {
			return false
		}
	}
	return true
}
