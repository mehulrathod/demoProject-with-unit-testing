package v1

type OrderMenuRequest struct {
	ID      uint    `form:"id"`
	OrderId uint    `form:"order_id" binding:"required"`
	MenuId  uint    `form:"menu_id" binding:"required"`
	Price   float64 `form:"price"`
}

type GetOrderMenuRequest struct {
	ID      uint    `form:"id"`
	OrderId uint `form:"order_id" binding:"required"`
}
