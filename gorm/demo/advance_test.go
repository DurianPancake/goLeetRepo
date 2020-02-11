package demo

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goSecond/gorm/demo/model"
	"log"
	"testing"
)

// initial database connection
func InitDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:easy@tcp(localhost:3306)/gorm_demo?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// 取消复数形式的表名
	db.SingularTable(true)
	return db
}

// create table
func TestCreateTable(t *testing.T) {
	db := InitDb()
	defer db.Close()

	db.SingularTable(true)
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.CreateTable(&model.User{})
	db.CreateTable(&model.Address{})
	db.CreateTable(&model.CreditCard{})
	db.CreateTable(&model.Email{})
	db.CreateTable(&model.Language{})
}
