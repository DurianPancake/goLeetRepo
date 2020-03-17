package fomular

type Stack struct {
	Kind    _type
	Field   string
	Operate Operator
	Value   string `json:"value"`
}

type _type uint8

const (
	Field _type = iota
	Operate
	Value
)

type Operator string

var operators = []Operator{
	eq,
	ne,
	gt,
	lt,
	gte,
	lte,
	and,
	or,
	//not
	like,
	in,
	notIn,
}

const (
	eq    Operator = "="
	ne    Operator = "!="
	gt    Operator = ">"
	lt    Operator = "<"
	gte   Operator = ">="
	lte   Operator = "<="
	and   Operator = "&&"
	or    Operator = "||"
	like  Operator = "like"
	in    Operator = "in"
	notIn Operator = "not_in"
	//not
)

func MatchOperator(char rune) (offset []int) {
	for _, str := range operators {
		firstChar := []rune(str)[0]
		if firstChar == char {
			offset = append(offset, len(str)-1)
		}
	}
	return
}

func Match(text string) Operator {
	for i, operator := range operators {
		if string(operator) == text {
			return operators[i]
		}
	}
	return ""
}
