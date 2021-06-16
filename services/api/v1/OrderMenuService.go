package v1Service

import (
	helpers "vipulsirdemo/apiHelpers"
	"vipulsirdemo/models"
	v1repo "vipulsirdemo/repository/api/v1"
	v1req "vipulsirdemo/resources/request/api/v1"
)

type OrderMenuService struct {
	OrderMenuRequest v1req.OrderMenuRequest
	OrderMenu        models.OrderMenu
	OrderMenuRepo    v1repo.OrderMenuRepo
}

func (oms *OrderMenuService) AddOrderMenu(omr v1req.OrderMenuRequest) map[string]interface{} {
	orderMenu := oms.OrderMenu
	orderMenu.OrderId = omr.OrderId
	orderMenu.MenuId = omr.MenuId
	userData := oms.OrderMenuRepo.AddOrderMenu(orderMenu)
	response := helpers.Message(helpers.ResponseSuccess, "Menu added in order.")
	response["data"] = userData
	return response
}

func (oms *OrderMenuService) GetOrderMenu(id uint) map[string]interface{} {
	res, err := oms.OrderMenuRepo.GetOrderMenus(id)
	if err != nil {
		response := helpers.Message(helpers.ResponseError, "something went wrong.")
		return response
	}
	response := helpers.Message(helpers.ResponseSuccess, "success.")
	response["data"] = res
	return response
}
