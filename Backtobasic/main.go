package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var stu []student

func main() {
	stu = append(stu, student{"surya", 88}, student{"raju", 99})

	r := gin.Default()
	r.GET("/getall", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, stu)
	})
	r.POST("/updatestudent", func(ctx *gin.Context) {
		var st1 student
		err := ctx.BindJSON(&st1)
		if err != nil {
			return
		}
		stu = append(stu, st1)
		ctx.IndentedJSON(http.StatusOK, stu)
	})
	r.PUT("/delete", func(ctx *gin.Context) {
		var st1 student
		err := ctx.BindJSON(&st1)
		if err != nil {
			return
		}
		stu = append(stu, st1)
		ctx.IndentedJSON(http.StatusOK, stu)
	})
	r.PUT("/delete", func(ctx *gin.Context) {
		var st1 student
		err := ctx.BindJSON(&st1)
		if err != nil {
			return
		}
		stu = append(stu, st1)
		ctx.IndentedJSON(http.StatusOK, stu)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
