package main

import (
	"goApi/controllers"
	_ "goApi/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	godotenv.Load()
	r := gin.Default()
	db := controllers.ConnectDB()

	//GET
	r.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.GET("/users/createtable", func(c *gin.Context) {
		controllers.CreateUsersTable(c, db)
	})
	r.GET("/users/get", func(c *gin.Context) {
		controllers.GetUsers(c, db)
	})
	r.GET("/user/get", func(c *gin.Context) {
		controllers.GetUserId(c, db)
	})
	//POST
	r.POST("/users/create", func(c *gin.Context) {
		controllers.CreateUsers(c, db)
	})
	r.POST("/users/post", func(c *gin.Context) {
		controllers.PostUsers(c, db)
	})
	r.POST("/user/post", func(c *gin.Context) {
		controllers.PostUserId(c, db)
	})
	//PATCH: actualización parcial de un registro. PUT: actualización completa.
	r.PATCH("/user/update", func(c *gin.Context) {
		controllers.UpdateUser(c, db)
	})
	//DELETE
	r.DELETE("/user/delete/:id", func(c *gin.Context) {
		controllers.DeleteUser(c, db)
	})

	r.Run()
}
