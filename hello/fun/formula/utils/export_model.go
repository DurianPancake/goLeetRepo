package utils

type Unit struct {
	Field    string      `json:"field_id"`
	Operator string      `json:"operate"`
	Value    interface{} `json:"value"`
	Type     string      `json:"type"`
}

type UnitGroup struct {
	Type       string        `json:"type"`
	Conditions []interface{} `json:"conditions"`
}

var logs = map[string]string{
	not.symbol: "not",
	and.symbol: "and",
	or.symbol:  "or",
}

const (
	condition = "condition"
)
