package v1

import (
	"github.com/jinzhu/gorm"
	"vipulsirdemo/models"
)

type HobbiesRepository interface {
	AddHobby(hobby models.Hobbies)
}

type HobbyRepo struct {
	DB *gorm.DB
}

func (hr *HobbyRepo) AddHobby(hobby models.Hobbies) {
	hr.DB.Model(&hobby).Create(&hobby)
}