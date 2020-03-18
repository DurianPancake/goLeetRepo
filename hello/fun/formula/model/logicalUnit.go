package model

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
