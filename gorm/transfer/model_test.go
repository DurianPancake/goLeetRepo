package transfer

import (
	"github.com/jinzhu/gorm"
	"log"
)

func initDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:easy@tcp(localhost:3306)/gorm_example?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// 取消复数形式的表名
	db.SingularTable(true)
	db.AutoMigrate(&User{})
	return db
}
