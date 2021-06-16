package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	helpers "vipulsirdemo/apiHelpers"
	v1req "vipulsirdemo/resources/request/api/v1"
	v1Service "vipulsirdemo/services/api/v1"
)

type MenuController struct {
	MenuService v1Service.MenuService
}

func (mc *MenuController) AddMenu(c *gin.Context) {
	menuService := mc.MenuService
	var menu v1req.AddEditMenuRequest
	if err := c.MustBindWith(&menu, binding.FormMultipart); err != nil {
		helpers.Respond(c.Writer, helpers.Message(helpers.ResponseError, "Invalid request."))
		return
	}

	ImageName, err := helpers.ImageUpload(c, "public/images/menus", "menu")
	if err {
		helpers.Respond(c.Writer, helpers.Message(helpers.ResponseError, "Upload valid file."))
		return
	}

	menu.ImageName = ImageName
	resp := menuService.AddMenu(menu)
	helpers.Respond(c.Writer, resp)
	return
}

func (mc *MenuController) EditMenu(c *gin.Context) {
	menuService := mc.MenuService
	var menu v1req.AddEditMenuRequest
	if err := c.MustBindWith(&menu, binding.FormMultipart); err != nil {
		helpers.Respond(c.Writer, helpers.Message(helpers.ResponseError, "Invalid request."))
		return
	}

	ImageName, err := helpers.ImageUpload(c, "public/images/menus", "menu")
	if err {
		helpers.Respond(c.Writer, helpers.Message(helpers.ResponseError, "Upload valid file."))
		return
	}

	menu.ImageName = ImageName
	resp := menuService.EditMenu(menu)
	helpers.Respond(c.Writer, resp)
	return
}

func (mc *MenuController) GetAllMenu(c *gin.Context) {
	menuService := mc.MenuService
	resp := menuService.GetAllMenu()
	helpers.Respond(c.Writer, resp)
	return
}

