package main

import (
	"goApi/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	db := controllers.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.GET("/users/create", func(c *gin.Context) {
		controllers.CreateUsers(c, db)
	})
	r.GET("/users/get", func(c *gin.Context) {
		controllers.GetUsers(c, db)
	})
	r.GET("/user/get", func(c *gin.Context) {
		controllers.GetUserId(c, db)
	})
	r.POST("/users/post", func(c *gin.Context) {
		controllers.PostUsers(c, db)
	})
	r.POST("/user/post", func(c *gin.Context) {
		controllers.PostUserId(c, db)
	})

	r.Run()
}
