package models

type Hobbies struct {
	Model
	Name string `gorm:"type:varchar(30)" json:"name" validate:"required"`
}

func (u *Hobbies) TableName() string {
	return "hobbies"
}
