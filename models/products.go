package models

type Products struct {
	Model
	Title string `gorm:"type:varchar(150)" json:"title" validate:"required"`
	Description string `gorm:"type:varchar(200)" json:"description"`
	Price float64 `json:"price" validate:"required"`
	Image string `gorm:"type:varchar(60)" json:"image"`
	Newsletter   uint   `gorm:"default:null" json:"newsletter"`
}

func (u *Products) TableName() string {
	return "products"
}

