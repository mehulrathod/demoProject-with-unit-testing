package models

type Order struct {
	Model
	UserId      uint    `json:"user_id" validate:"required"`
	OrderNumber string  `gorm:"type:varchar(10)" json:"order_number" validate:"required"`
	Description string  `gorm:"type:varchar(200)" json:"description"`
	//Date        time.Time `json:"date" validate:"required"`
	Status      string  `gorm:"type:enum('active', 'inactive')" json:"status"`
}

func (u *Order) TableName() string {
	return "orders"
}
