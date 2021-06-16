package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	helpers "vipulsirdemo/apiHelpers"
	v1req "vipulsirdemo/resources/request/api/v1"
	v1Service "vipulsirdemo/services/api/v1"
)

type ProductController struct {
	ProductService v1Service.ProductService
}

func (pc *ProductController) AddProduct(c *gin.Context) {
	productService := pc.ProductService
	var product v1req.AddProductRequest
	if err := c.MustBindWith(&product, binding.FormMultipart); err != nil {
		helpers.Respond(c.Writer, helpers.Message(helpers.ResponseError, "Invalid request."))
		return
	}

	ImageName, err := helpers.ImageUpload(c, "public/images/products", "product")
	if err {
		helpers.Respond(c.Writer, helpers.Message(helpers.ResponseError, "Upload valid file."))
		return
	}
	product.ImageName = ImageName

	resp := productService.AddProduct(product)
	helpers.Respond(c.Writer, resp)
	return
}