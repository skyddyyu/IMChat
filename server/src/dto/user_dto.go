// dto/user_dto.go
package dto

import "time"

type UserRegisterDTO struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserLoginDTO struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserLoginResponseDTO struct {
	ID             string `json:"id"`
	Account        string `json:"account"`
	Token          string `json:"token"`
	ProfilePicture string `json:"profile_picture"`
	// 其他字段...
}

type UserResponseDTO struct {
	ID             string     `json:"id"`
	Account        string     `json:"account"`
	Gender         string     `json:"gender"`
	Bio            string     `json:"bio"`
	ProfilePicture string     `json:"profile_picture"`
	LastLogin      *time.Time `json:"last_login"`
}

type ChatRoomUserListResponseDTO struct {
	ID             string     `json:"id"`
	Account        string     `json:"account"`
	Gender         string     `json:"gender"`
	Bio            string     `json:"bio"`
	ProfilePicture string     `json:"profile_picture"`
	LastLogin      *time.Time `json:"last_login"`
	Active         bool       `json:"active"`
	IsAdmin        bool       `json:"is_admin"`
}
