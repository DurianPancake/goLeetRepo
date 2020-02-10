package gorm

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique_index"`
	Email    string
	Age      int
	Birthday time.Time
}
