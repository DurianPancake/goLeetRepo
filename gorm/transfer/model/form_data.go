package model

type FormData struct {
	TenantModel
	FormID         string `json:"form_id"`
	Type           string `json:"type"`
	Creator        string `json:"creator"`
	Editor         string `json:"editor"`
	Index          int    `json:"index"`
	ApprovalSt     string `json:"approval_st"`
	ApprovalStatus string `json:"approval_status"`
	FormSource     string `json:"form_source"`
	FormSources    string `json:"form_sources"`
	FormStatus     string `json:"form_status"`
	RuleIds        string `json:"rule_ids"`
	Status         string `json:"status"`
	TreeId         string `json:"tree_id"`
	TreeNodeID     string `json:"tree_node_id"`
	Data           []Data `json:"data" gorm:"ForeignKey:FormDataID"`
}

type Data struct {
	ID         string `json:"id"`           // 该字段表的主键
	FormDataID string `json:"form_data_id"` // 外键，关联到formData
	ParentID   string `json:"parent_id"`    // 明细字段时，用来表示该字段属于哪一个list，与ID形成树形结构
	RowID      string `json:"row_id"`       // 明细字段时，用来表示该字段属于哪一个row，业务中的rowId
	FieldName  string `json:"field_name"`
	Type       string `json:"type"`
	Value      string `json:"value"` // 当Type为list时，Value为空
}
