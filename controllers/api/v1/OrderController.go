package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	helper "vipulsirdemo/apiHelpers"
	v1req "vipulsirdemo/resources/request/api/v1"
	v1Service "vipulsirdemo/services/api/v1"
)

type OrderController struct {
	OrderService v1Service.OrderService
}

func (oc *OrderController) AddOrder(c *gin.Context) {
	orderService := oc.OrderService
	var order v1req.OrderRequest
	if err := c.MustBindWith(&order, binding.FormMultipart); err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
		return
	}
	resp := orderService.AddOrder(order)
	helper.Respond(c.Writer, resp)
	return
}

func (oc *OrderController) GetOrders(c *gin.Context) {
	orderService := oc.OrderService
	var order v1req.GetUserOrderRequest
	if err := c.MustBindWith(&order, binding.FormMultipart); err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
		return
	}
	resp := orderService.GetOrders(order.UserId)
	helper.Respond(c.Writer, resp)
	return
}
