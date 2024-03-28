// models/user.go

package models

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserIDRequest struct {
	UserID uint `json:"userId"`
}
