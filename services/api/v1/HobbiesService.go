package v1Service

import (
	helper "vipulsirdemo/apiHelpers"
	"vipulsirdemo/models"
	v1repo "vipulsirdemo/repository/api/v1"
	v1res "vipulsirdemo/resources/response/api/v1"
)

type HobbiesService struct {
	Hobby     models.Hobbies
	HobbyRepo v1repo.HobbyRepo
}

func (hs *HobbiesService) AddHobby() map[string]interface{} {
	hobby := hs.Hobby
	hs.HobbyRepo.AddHobby(hobby)
	hobbyData := v1res.HobbyResponse{
		Name: hobby.Name,
	}
	response := helper.Message(helper.ResponseSuccess, "Hobby added successfully.")
	response["data"] = hobbyData
	return response
}
