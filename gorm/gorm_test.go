package gorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"testing"
	"time"
)

func initDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:easy@tcp(localhost:3306)/gorm_example?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// 取消复数形式的表名
	db.SingularTable(true)
	return db
}

func TestConnect(t *testing.T) {
	db := initDb()
	fmt.Printf("type: %T, value: %v", db, db)
}

func TestInsert(t *testing.T) {
	db := initDb()
	defer db.Close()

	user := &User{
		Username: "Jack Ma",
		Email:    "1@1.com",
		Age:      18,
	}
	create := db.Create(user)
	fmt.Println(create.GetErrors())
}

func TestUpdate(t *testing.T) {
	db := initDb()
	defer db.Close()

	//db.Model(&User{}).Where("age<?", 20).Update("username", "Tony Ma")
	db.Model(&User{}).Where("id=?", 5).Update(map[string]interface{}{
		"username": "Tony ba ba",
		"email":    "2@2.com",
		"age":      20,
		"birthday": time.Date(2020, 2, 2, 0, 0, 0, 0, time.UTC),
	})
	fmt.Println(db.GetErrors())
}

func TestUpdate2(t *testing.T) {
	db := initDb()
	defer db.Close()

	db.Model(&User{}).Where("birthday=?", "0000-00-00").Update("birthday", "2020-02-10")
	fmt.Println(db.GetErrors())
}

func TestDelete(t *testing.T) {
	db := initDb()
	defer db.Close()

	db.Delete(&User{}).Where("id=?", 3)
}

func TestQuery(t *testing.T) {
	db := initDb()
	defer db.Close()

	user := User{}
	// 查询第一个
	db.First(&user)
	fmt.Println(user)
	// 查询所有数据
	var users []User
	db.Find(&users)
	fmt.Println(users)

	// 模糊查询
	db.Where("username like ?", "%ba%").Find(&users)
	fmt.Println(users)
}
