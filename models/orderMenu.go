package models

import v1Response "vipulsirdemo/resources/response/api/v1"

type OrderMenu struct {
	Model
	OrderId uint                            `json:"order_id" validate:"required"`
	MenuId  uint                            `json:"menu_id" validate:"required"`
	Price   float64                         `json:"price" validate:"required"`
	Menu    v1Response.CategoryMenuResponse `json:"menu"`
}

func (u *OrderMenu) TableName() string {
	return "order_menu"
}
