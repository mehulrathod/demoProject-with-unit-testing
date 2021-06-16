package v1Response

//import "time"

type UserResponse struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Image string `json:"image,omitempty"`
	//Dob time.Time	`json:"dob,omitempty"`
	Gender string `json:"gender,omitempty"`
	Mobile int64  `json:"mobile,omitempty"`
	Hobby  uint   `json:"hobby"`
}

type LoginResponse struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Image string `json:"image,omitempty"`
	//Dob time.Time	`json:"dob,omitempty"`
	Gender string `json:"gender,omitempty"`
	Mobile int64  `json:"mobile,omitempty"`
	Hobby  uint   `json:"hobby"`
	Token  string `json:"token"`
}
