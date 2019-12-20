package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)
type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"username" json:"password" binding:"required"`
}
func main(){
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping",func(c *gin.Context){
		c.JSON(200,gin.H{
			"msg":"1pong",
		})
	})
	r.POST("/login",func(c *gin.Context){
		var u User
		if err:= c.BindJSON(&u); err!=nil{
			c.JSON(200,gin.H{
				"msg":"param error",
			})
			return
		}
		if u.Username == "123" && u.Password == "456" {
			c.JSON(200,gin.H{
				"msg":"success",
			})
		} else{
			c.JSON(200,gin.H{
				"msg":"wrong",
			})
		}
	})
	r.Run(":3000")
}

