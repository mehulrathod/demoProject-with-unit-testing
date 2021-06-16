package v1

import (
	"time"
)

type OrderRequest struct {
	ID          uint
	UserId      uint      `form:"user_id" binding:"required"`
	OrderNumber string    `form:"order_number" binding:"required"`
	Description string    `form:"description"`
	Date        time.Time `form:"date" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	Status      string    `form:"status"`
}

type GetUserOrderRequest struct {
	ID     uint
	UserId uint `form:"user_id" binding:"required"`
}
