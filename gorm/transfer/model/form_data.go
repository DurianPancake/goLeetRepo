package model

type formData struct {
	TenantModel
	FormID         string `json:"form_id"`
	Type           string `json:"type"`
	Creator        string
	Editor         string
	Index          int
	ApprovalSt     string `json:"approval_st"`
	ApprovalStatus string `json:"approval_status"`
	FormSource     string `json:"form_source"`
	FormSources    string `json:"form_sources"`
	FormStatus     string `json:"form_status"`
	RuleIds        string `json:"rule_ids"`
	Status         string
	TreeId         string `json:"tree_id"`
	TreeNodeID     string `json:"tree_node_id"`
	Data           []Data `json:"data"`
}

type Data struct {
	FieldName string `json:"field_name"`
	Type      string
}
