// controllers/user.go

package controllers

import (
	"goApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
	if err := c.ShouldBindQuery(&userIDRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if err := db.First(&user, userIDRequest.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func PostUsers(c *gin.Context, db *gorm.DB) {
	var users []models.User
	db.Find(&users)
	// db.Where(models.User{Name: "pp"}).Find(&users)
	c.JSON(http.StatusOK, users)
}

func PostUserId(c *gin.Context, db *gorm.DB) {
	var userIDRequest models.UserIDRequest
	var user models.User
	if err := c.ShouldBindJSON(&userIDRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if err := db.First(&user, userIDRequest.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
