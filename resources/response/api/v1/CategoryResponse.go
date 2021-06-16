package v1Response

type CategoryResponse struct {
	Id       uint   `json:"id,"`
	Name     string `json:"name,omitempty"`
	Status   string  `json:"status"`
	Menu []CategoryMenuResponse `json:"menu"`
}
