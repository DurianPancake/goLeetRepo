package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Profile   Profile
	ProfileId int
}

type Profile struct {
	gorm.Model
	Name string
}
