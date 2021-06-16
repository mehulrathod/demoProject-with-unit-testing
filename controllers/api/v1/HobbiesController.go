package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	helper "vipulsirdemo/apiHelpers"
	v1Service "vipulsirdemo/services/api/v1"
)

type HobbiesController struct {
	HobbiesService v1Service.HobbiesService
}

func (hc *HobbiesController) AddHobby(c *gin.Context) {
	hobbyService := hc.HobbiesService
	err := json.NewDecoder(c.Request.Body).Decode(&hobbyService.Hobby)
	if err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
	}
	resp := hobbyService.AddHobby()
	helper.Respond(c.Writer, resp)
	return
}
