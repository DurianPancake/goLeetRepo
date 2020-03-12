package model

type FieldItem struct {
	Type      string      `json:"type"`
	Operate   string      `json:"operate"`
	Value     interface{} `json:"value"`
	ShowValue interface{} `json:"show_value"`
}
