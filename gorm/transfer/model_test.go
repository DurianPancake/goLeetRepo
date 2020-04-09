package transfer

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"goSecond/gorm/transfer/model"
	"log"
	"testing"
)

func initDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:easy@tcp(localhost:3306)/gorm_example?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// 取消复数形式的表名
	db.SingularTable(true)
	db.AutoMigrate(&model.FormData{})
	db.AutoMigrate(&model.Data{})
	return db
}

func TestT(t *testing.T) {
	initDb()
}
