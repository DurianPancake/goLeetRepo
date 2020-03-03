package related

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goSecond/gorm/related/model"
	"log"
	"testing"
)

// initial database connection
func InitDb() *gorm.DB {
	db, err := gorm.Open("mysql",
		"root:easy@tcp(localhost:3306)/gorm_example?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// 取消复数形式的表名
	db.SingularTable(true)
	db.AutoMigrate(&model.User{}, &model.Profile{})
	return db
}

func TestRe1(t *testing.T) {
	db := InitDb()
	defer db.Close()

	user := model.User{
		Profile: model.Profile{
			Name: "HEhe",
		},
		ProfileId: 1,
	}
	db.Create(&user)
	fmt.Println(db.GetErrors())
}

func TestReQuery(t *testing.T) {
	db := InitDb()

	defer db.Close()
	var user model.User
	db.Model(&model.User{}).Related(&model.Profile{}).First(&user)

	fmt.Println(user)
	fmt.Println(db.GetErrors())
}
