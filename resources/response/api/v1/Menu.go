package v1Response

type MenuResponse struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description"`
	Price       float64 `json:"price,omitempty"`
	Image       string  `json:"image,omitempty"`
	Status      string  `json:"status"`
	Category    string  `json:"category"`
}

type CategoryMenuResponse struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description"`
	Price       float64 `json:"price,omitempty"`
	Image       string  `json:"image,omitempty"`
}
