package model

import (
	"time"
)

//数据库基础model
type Model struct {
	ID        string     `gorm:"primary_key;auto_increment:false" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

//数据库基础model with tenant and by
type TenantModel struct {
	Model
	Tenant string `json:"tenant,omitempty" gorm:"primary_key;auto_increment:false"`
}
