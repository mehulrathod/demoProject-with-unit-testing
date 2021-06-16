package v1Response

type OrderMenuResponse struct {
	Id      uint                 `json:"id"`
	OrderId uint                 `json:"order_id,omitempty"`
	MenuId  uint                 `json:"menu_id,omitempty"`
	Menu    CategoryMenuResponse `json:"menu"`
	Price   float64              `json:"price,omitempty"`
}

type OrdersOrderMenuResponse struct {
	Id    uint                 `json:"id"`
	Menu  CategoryMenuResponse `json:"menu"`
	Price float64              `json:"price,omitempty"`
}

type AddOrderMenuResponse struct {
	MenuId uint  `json:"menu_id"`
	Price  float64 `json:"price,omitempty"`
}
