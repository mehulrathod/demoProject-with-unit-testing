package v1

import (
	"mime/multipart"
	"time"
)

type SignUpRequest struct {
	ID        uint
	Name string    `form:"name" binding:"required"`
	Email     string    `form:"email" binding:"required,email"`
	Mobile    int64    `form:"mobile" binding:"required"`
	Image     *multipart.FileHeader `form:"image" binding:"required"`
	ImageName string
	Password  string    `form:"password" binding:"required"`
	Hobby uint
	Gender string    `form:"gender" binding:"required"`
	CreatedAt time.Time `form:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `form:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}

type LoginRequest struct {
	Email     string    `form:"email" binding:"required,email"`
	Password  string    `form:"password" binding:"required"`
}

type EditProfileRequest struct {
	Id uint
	Name string    `form:"name" binding:"required"`
	Mobile    int64    `form:"mobile" binding:"required"`
	Image     *multipart.FileHeader `form:"image" binding:"required"`
	ImageName string
	Hobby uint
	Gender string    `form:"gender" binding:"required"`
}
