package v1Service

import (
	"fmt"
	helpers "vipulsirdemo/apiHelpers"
	"vipulsirdemo/models"
	v1repo "vipulsirdemo/repository/api/v1"
	v1req "vipulsirdemo/resources/request/api/v1"
	v1res "vipulsirdemo/resources/response/api/v1"
)

type ProductService struct {
	ProductRequest v1req.AddProductRequest
	Product        models.Products
	ProductRepo    v1repo.ProductRepo
}

func (ps *ProductService) AddProduct(pr v1req.AddProductRequest) map[string]interface{} {
	product := ps.Product
	fmt.Println("us", pr)
	product.Title = pr.Title
	product.Description = pr.Description
	product.Price = pr.Price
	product.Image = pr.ImageName
	ps.ProductRepo.AddProduct(product)
	productData := v1res.ProductResponse{
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Image:       "http://localhost:" + helpers.GetPort() + "/public/images/product/" + product.Image,
	}
	response := helpers.Message(helpers.ResponseSuccess, "Product added successfully.")
	response["data"] = productData
	return response
}
