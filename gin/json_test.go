package mygin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goSecond/gin/model"
	"log"
	"testing"
)

func TestPostJson(t *testing.T) {

	engine := gin.Default()

	engine.POST(" /student", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var stu Student
		if err := context.BindJSON(&stu); err != nil {
			log.Fatal(err)
		}
		fmt.Println("name:", stu.name)
		fmt.Println("age:", stu.age)
		_, _ = context.Writer.WriteString("success Register")
	})

	engine.Run()
}

type Student struct {
	name string
	age  int
}

func TestReturnJSON(t *testing.T) {
	engine := gin.Default()

	engine.GET("/helloJson", func(context *gin.Context) {
		fullPath := "请求路径" + context.FullPath()
		fmt.Println(fullPath)

		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  "ok",
			"data": fullPath,
		})
	})

	engine.GET("/jsonStruct", func(context *gin.Context) {
		fullPath := "请求路径: " + context.FullPath()
		fmt.Println(fullPath)

		resp := model.Success("ok", fullPath)
		context.JSON(200, resp)
	})

	engine.Run(":8091")
}
