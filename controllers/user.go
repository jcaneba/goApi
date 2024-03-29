// controllers/user.go

package controllers

import (
	"goApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

func CreateUsers(c *gin.Context, db *gorm.DB) {
	db.AutoMigrate(models.User{})
}

func GetUsers(c *gin.Context, db *gorm.DB) {
	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserId(c *gin.Context, db *gorm.DB) {
	var userIDRequest models.UserIDRequest
	var user models.User
	if err := c.ShouldBindQuery(&userIDRequest); err != nil { //Para UserIDRequest con form:"userId"
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&user, userIDRequest.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func PostUsers(c *gin.Context, db *gorm.DB) {
	var users []models.User
	db.Find(&users)
	// db.Where(models.User{Name: "pp"}).Find(&users)
	// db.Where("name LIKE ?", "%pp%").Find(&users)
	// db.Where(&User{Age: gorm.LT(25)}).Find(&users)
	c.JSON(http.StatusOK, users)
}

func PostUserId(c *gin.Context, db *gorm.DB) {
	var userIDRequest models.UserIDRequest
	var user models.User
	if err := c.ShouldBindWith(&userIDRequest, binding.Form); err != nil { //Para UserIDRequest con form:"userId"
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} /*else if err := c.ShouldBindJSON(&userIDRequest); err != nil { //Para UserIDRequest con json:"userId"
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}*/

	if err := db.First(&user, userIDRequest.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
