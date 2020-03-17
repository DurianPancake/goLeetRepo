package fomular

type Stack struct {
	Kind      _type
	Condition string
	Operate   Operator
	Value     string `json:"value"`
}

type _type uint8

const (
	Condition _type = iota
	Operate
	Value
)

type Operator string

const (
	eq    = "="
	gt    = ">"
	lt    = "<"
	gte   = ">="
	lte   = "<="
	and   = "&&"
	or    = "||"
	not   = "!"
	in    = "in"
	notIn = "not in"
)
