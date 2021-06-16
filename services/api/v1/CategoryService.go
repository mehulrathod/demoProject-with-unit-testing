package v1Service

import (
	"fmt"
	helpers "vipulsirdemo/apiHelpers"
	"vipulsirdemo/models"
	v1repo "vipulsirdemo/repository/api/v1"
	v1req "vipulsirdemo/resources/request/api/v1"
	v1res "vipulsirdemo/resources/response/api/v1"
)

type CategoryService struct {
	CategoryRequest v1req.AddEditCategoryRequest
	Category        models.Category
	CategoryRepo    v1repo.CategoryRepo
}

func (cs *CategoryService) AddCategory(cr v1req.AddEditCategoryRequest) map[string]interface{} {
	category := cs.Category
	category.Name = cr.Name
	category.Status = "active"
	cs.CategoryRepo.AddCategory(category)
	userData := v1res.CategoryResponse{
		Name: category.Name,
	}
	response := helpers.Message(helpers.ResponseSuccess, "Category added successfully.")
	response["data"] = userData
	return response
}

func (cs *CategoryService) EditCategory(cr v1req.AddEditCategoryRequest) map[string]interface{} {
	category := cs.Category
	category.ID = cr.ID
	category.Name = cr.Name
	category.Status = "active"
	res, err := cs.CategoryRepo.GetCategoryById(category.ID)
	if err != nil {
		fmt.Println(res)
		return helpers.Message(helpers.ResponseError, "Enter valid category.")
	}
	cs.CategoryRepo.EditCategory(category)
	userData := v1res.CategoryResponse{
		Name: category.Name,
	}
	response := helpers.Message(helpers.ResponseSuccess, "Category updated successfully.")
	response["data"] = userData
	return response
}

func (cs *CategoryService) GetAllCategories() map[string]interface{} {
	res, err := cs.CategoryRepo.GetAllCategories()
	if err != nil {
		response := helpers.Message(helpers.ResponseError, "something went wrong.")
		return response
	}
	fmt.Println("all categories", res, err)
	fmt.Println(cs.CategoryRepo.GetCategoryFromRedis())
	response := helpers.Message(helpers.ResponseSuccess, "success.")
	response["data"] = res
	return response
}
