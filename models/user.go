// models/user.go

package models

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

/*
//Para el endpoint "PostUserId" pasando el JSON {"userId":1} como parámetro POST

	type UserIDRequest struct {
		UserID uint `json:"userId"`
	}
*/
type UserIDRequest struct {
	UserID uint `form:"userId" binding:"required"`
}
