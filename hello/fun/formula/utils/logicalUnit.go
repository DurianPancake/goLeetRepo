package utils

// 运算符
type Operator struct {
	Symbol string
	Type   string
}

// 基本逻辑运算单元，代表一个布尔值类型
type Unit struct {
	Field   string
	Operate Operator
	Value   string `json:"value"`
}

// 描述一组逻辑运算关系，在该组中，逻辑运算符是一样的，表示同一优先级的运算
// 基本单位可以是Unit或是Logical
// ! > && > ||
// ! 只能跟一个Unit或LogicalGroup
type LogicalGroup struct {
	Operator
	Units        []Unit
	LogicalUnits []LogicalGroup
}

var operators = []Operator{
	Eq,
	Ne,
	Gt,
	Lt,
	Gte,
	Lte,
	Like,
	In,
	NotIn,
	And,
	Or,
	Not,
	LeftBracket,
	RightBracket,
	Null,
}

const (
	base    = "base"
	logic   = "logic"
	bracket = "bracket"
)

var (
	Eq    = Operator{"=", base}
	Ne    = Operator{"!=", base}
	Gt    = Operator{">", base}
	Lt    = Operator{"<", base}
	Gte   = Operator{">=", base}
	Lte   = Operator{"<=", base}
	Like  = Operator{"like", base}
	In    = Operator{"in", base}
	NotIn = Operator{"not_in", base}
	//
	Not = Operator{"!", logic}
	And = Operator{"&&", logic}
	Or  = Operator{"||", logic}
	//
	LeftBracket  = Operator{"(", bracket}
	RightBracket = Operator{")", bracket}
	//
	Null = Operator{"", ""}
)

// 最大匹配原则，尽可能长的匹配字段或符号
// 获取每个匹配的偏移量，选取偏移量最大的
func MatchOperate(text string, index int) Operator {

	runes := []rune(text)
	maxOffset := 0
	var chosenOp Operator
	for _, operator := range operators {
		if operator == Null {
			continue
		}
		firstChar := []rune(operator.Symbol)[0]
		if firstChar == runes[index] {
			if length := len(operator.Symbol); length > maxOffset && string(runes[index:index+length]) == operator.Symbol {
				maxOffset = length
				chosenOp = operator
			}
		}
	}
	return chosenOp
}

// 寻找匹配右括号的索引位置
// ((() ())())
func FindMatchBracketIndex(text string, index int) int {

	operate := MatchOperate(text, index)
	if operate.Type != bracket || operate.Symbol != "(" {
		return -1
	}
	//
	count := 1
	runes := []rune(text)
	for i := index + 1; i < len(runes); i++ {
		if runes[i] == '(' {
			count += 1
		}
		if runes[i] == ')' {
			count -= 1
			if count == 0 {
				return i
			}
		}
	}
	return -1
}

func (op1 *Operator) equals(op2 *Operator) bool {
	return op1.Symbol == op2.Symbol
}

func (u1 *Unit) equals(u2 *Unit) bool {
	return u1.Operate.equals(&u2.Operate) && u1.Field == u2.Field && u1.Value == u2.Value
}

// 判断逻辑组是否相等
func (lg1 *LogicalGroup) equals(lg2 *LogicalGroup) bool {
	return lg1.Symbol == lg2.Symbol && isUnitsEquals(lg1, lg2) && isLogicalUnitsEquals(lg1, lg2)
}

func isLogicalUnitsEquals(lg1 *LogicalGroup, lg2 *LogicalGroup) bool {
	if len(lg1.LogicalUnits) != len(lg2.LogicalUnits) {
		return false
	}
	for i, unit := range lg1.LogicalUnits {
		units := lg2.LogicalUnits
		if !units[i].equals(&unit) {
			return false
		}
	}
	return true
}

func isUnitsEquals(lg1 *LogicalGroup, lg2 *LogicalGroup) bool {
	if len(lg1.Units) != len(lg2.Units) {
		return false
	}
	for i, unit := range lg1.Units {
		units := lg2.Units
		if !units[i].equals(&unit) {
			return false
		}
	}
	return true
}
