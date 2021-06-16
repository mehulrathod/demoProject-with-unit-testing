package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	helper "vipulsirdemo/apiHelpers"
	v1req "vipulsirdemo/resources/request/api/v1"
	v1Service "vipulsirdemo/services/api/v1"
)

type CategoryController struct {
	CategoryService v1Service.CategoryService
}

func (cc *CategoryController) AddCategory(c *gin.Context) {
	categoryService := cc.CategoryService
	var category v1req.AddEditCategoryRequest
	err := json.NewDecoder(c.Request.Body).Decode(&category)
	if err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
		return
	}
	resp := categoryService.AddCategory(category)
	helper.Respond(c.Writer, resp)
	return
}

func (cc *CategoryController) EditCategory(c *gin.Context) {
	categoryService := cc.CategoryService
	var category v1req.AddEditCategoryRequest
	err := json.NewDecoder(c.Request.Body).Decode(&category)
	if err != nil {
		helper.Respond(c.Writer, helper.Message(helper.ResponseError, "Invalid request."))
		return
	}
	resp := categoryService.EditCategory(category)
	helper.Respond(c.Writer, resp)
	return
}

func (cc *CategoryController) GetAllCategories(c *gin.Context) {
	categoryService := cc.CategoryService
	resp := categoryService.GetAllCategories()
	helper.Respond(c.Writer, resp)
	return
}
