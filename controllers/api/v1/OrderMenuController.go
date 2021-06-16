package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	helper "vipulsirdemo/apiHelpers"
	v1req "vipulsirdemo/resources/request/api/v1"
	v1Service "vipulsirdemo/services/api/v1"
)

type OrderMenuController struct {
	OrderMenuService v1Service.OrderMenuService
}

func (omc *OrderMenuController) AddOrderMenu(c *gin.Context) {
	orderMenuService := omc.OrderMenuService
	var orderMenu v1req.OrderMenuRequest
	if err := c.MustBindWith(&orderMenu, binding.FormMultipart); err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
		return
	}
	fmt.Println("con ", orderMenu, orderMenu.OrderId)
	resp := orderMenuService.AddOrderMenu(orderMenu)
	helper.Respond(c.Writer, resp)
	return
}

func (omc *OrderMenuController) GetOrderMenu(c *gin.Context) {
	orderMenuService := omc.OrderMenuService
	var orderMenu v1req.GetOrderMenuRequest
	if err := c.MustBindWith(&orderMenu, binding.FormMultipart); err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
		return
	}
	resp := orderMenuService.GetOrderMenu(orderMenu.OrderId)
	helper.Respond(c.Writer, resp)
	return
}
