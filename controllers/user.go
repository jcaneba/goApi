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

// CreateUsers		godoc
// @Summary			Creación de usuarios
// @Description		Endpoint para crear un nuevo registro en la tabla "users"
// @Tags			Usuarios
// @Accept			application/json
// @Produce			application/json
// @Param			input body []models.User true "Información del usuario a crear"
// @Success			200 {object} []models.User
// @Router			/users/create [post]
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

// GetUsers			godoc
// @Summary			Obtención de usuarios
// @Description		Endpoint para recoger todos los datos de los usuarios en la tabla "users"
// @Tags			Usuarios
// @Produce			application/json
// @Success			200 {object} []models.User
// @Router			/users/get [get]
func GetUsers(c *gin.Context, db *gorm.DB) {
	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GetUserId		godoc
// @Summary			Obtención de un usuario
// @Description		Endpoint para recoger todos los datos de un único usuario en la tabla "users" a través de su ID
// @Tags			Usuarios
// @Accept			multipart/form-data
// @Produce			application/json
// @Success			200 {object} models.User
// @Router			/user/get [get]
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

// PostUsers		godoc
// @Summary			Obtención de usuarios
// @Description		Endpoint para recoger todos los datos de los usuarios en la tabla "users"
// @Tags			Usuarios
// @Produce			application/json
// @Success			200 {object} []models.User
// @Router			/users/post [post]
func PostUsers(c *gin.Context, db *gorm.DB) {
	var users []models.User
	db.Find(&users)
	// db.Where(models.User{Name: "pp"}).Find(&users)
	// db.Where("name LIKE ?", "%pp%").Find(&users)
	// db.Where(&User{Age: gorm.LT(25)}).Find(&users)
	c.JSON(http.StatusOK, users)
}

// PostUserId		godoc
// @Summary			Obtención de un usuario
// @Description		Endpoint para recoger todos los datos de un único usuario en la tabla "users" a través de su ID
// @Tags			Usuarios
// @Accept			multipart/form-data
// @Produce			application/json
// @Success			200 {object} models.User
// @Router			/user/post [post]
func PostUserId(c *gin.Context, db *gorm.DB) {
	var userIDRequest models.UserIDRequest
	var user models.User
	if err := c.ShouldBindWith(&userIDRequest, binding.Form); err != nil { //Para UserIDRequest con form:"userId"
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	/*if err := c.ShouldBindJSON(&userIDRequest); err != nil { //Para UserIDRequest con json:"userId"
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}*/

	if err := db.First(&user, userIDRequest.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser		godoc
// @Summary			Actualización de un usuario
// @Description		Endpoint para actualizar los datos de un usuario en la tabla "users"
// @Tags			Usuarios
// @Accept			application/json
// @Produce			application/json
// @Success			200 {object} models.User
// @Router			/user/update [patch]
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

// DeleteUser		godoc
// @Summary			Borrado de un usuario
// @Description		Endpoint para eliminar por completo un usuario en la tabla "users" a través de su ID
// @Tags			Usuarios
// @Produce			application/json
// @Success			204 "No Content"
// @Success			200 {object} map[string]string
// @Router			/user/delete/{id} [delete]
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
