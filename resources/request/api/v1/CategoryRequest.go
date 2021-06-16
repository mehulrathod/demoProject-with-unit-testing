package v1

import (
	"time"
)

type AddEditCategoryRequest struct {
	ID        uint
	Name      string    `form:"name" binding:"required"`
	Status    string    `form:"gender" binding:"status"`
	CreatedAt time.Time `form:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
