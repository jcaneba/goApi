// controllers/user.go

package controllers

import (
	"goApi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

func CreateUsersTable(c *gin.Context, db *gorm.DB) {
	db.AutoMigrate(models.User{})
}

func CreateUsers(c *gin.Context, db *gorm.DB) {
	var users []models.User
	if err := c.BindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
			return
		}
	}
	c.JSON(http.StatusCreated, users)
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context, db *gorm.DB) {
	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar si el usuario existe
	var existingUser models.User
	if err := db.First(&existingUser, updatedUser.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Actualizar el usuario en la base de datos
	if err := db.Model(&existingUser).Updates(&updatedUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el usuario"})
		return
	}

	c.JSON(http.StatusOK, existingUser)
}

func DeleteUser(c *gin.Context, db *gorm.DB) {
	var user models.User
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64) //De String a Int, en base 10 y tamaño de bits 64
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	// Eliminar el usuario de la base de datos
	if err := db.Delete(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente"})
}
