package v1Service

import (
	helpers "vipulsirdemo/apiHelpers"
	"vipulsirdemo/models"
	v1repo "vipulsirdemo/repository/api/v1"
	v1req "vipulsirdemo/resources/request/api/v1"
	v1res "vipulsirdemo/resources/response/api/v1"
)

type OrderService struct {
	OrderRequest v1req.OrderRequest
	Order        models.Order
	OrderRepo    v1repo.OrderRepo
}

func (ors *OrderService) AddOrder(or v1req.OrderRequest) map[string]interface{} {
	order := ors.Order
	order.OrderNumber = or.OrderNumber
	order.Description = or.Description
	//order.Date = or.Date
	order.Status = or.Status
	order.UserId = or.UserId
	ors.OrderRepo.AddOrder(order)
	userData := v1res.OrderResponse{
		OrderNumber: order.OrderNumber,
		Description: order.Description,
		//Date: order.Date,
		Status: order.Status,
		UserId: order.UserId,

	}
	response := helpers.Message(helpers.ResponseSuccess, "Order added successfully.")
	response["data"] = userData
	return response
}

func (ors *OrderService) GetOrders(id uint) map[string]interface{} {
	res, err := ors.OrderRepo.GetOrderByUser(id)
	if err != nil {
		response := helpers.Message(helpers.ResponseError, "something went wrong.")
		return response
	}
	response := helpers.Message(helpers.ResponseSuccess, "success.")
	response["data"] = res
	return response
}
