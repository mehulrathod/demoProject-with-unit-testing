package v1Response

type ProductResponse struct {
	Title string `json:"title,omitempty"`
	Description string `json:"description"`
	Price float64 `json:"price,omitempty"`
	Image string `json:"image,omitempty"`
}
