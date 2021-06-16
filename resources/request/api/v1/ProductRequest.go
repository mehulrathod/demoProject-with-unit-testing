package v1

import (
	"mime/multipart"
	"time"
)

type AddProductRequest struct {
	ID        uint
	Title string    `form:"title" binding:"required"`
	Description     string    `form:"description"`
	Price    float64    `form:"price" binding:"required"`
	Image     *multipart.FileHeader `form:"image" binding:"required"`
	ImageName string
	Newsletter uint `form:"newsletter"`
	CreatedAt time.Time `form:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `form:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
