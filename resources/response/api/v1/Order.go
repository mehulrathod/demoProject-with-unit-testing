package v1Response


type OrderResponse struct {
	Id          uint   `json:"id,"`
	UserId      uint   `json:"user_id,omitempty"`
	OrderNumber string `json:"order_number,omitempty"`
	Description string `json:"description"`
	//Date        time.Time `json:"date,omitempty"`
	Status string  `json:"status"`
	Total  string  `json:"total"`
	Items  []OrderMenuResponse `json:"items"`
}
