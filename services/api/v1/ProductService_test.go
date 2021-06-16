package v1Service

import (
	"github.com/stretchr/testify/mock"
	"testing"
	helpers "vipulsirdemo/apiHelpers"
	"vipulsirdemo/mocks"
	"vipulsirdemo/models"
	v1req "vipulsirdemo/resources/request/api/v1"
)

var ImageName = "t-shirt.jpeg"
var Image = "http://localhost:" + helpers.GetPort() + "/public/images/product/" + ImageName


var ProductData = &models.Products{
	Model: models.Model{
		ID: 1,
	},
	Title: "testing",
	Description: "its a testing description",
	Price: 20.50,
	Image: Image,
	Newsletter: 1,
}

var ProductRequestData = v1req.AddProductRequest{
	ID: 1,
	Title: "Testing",
	Description: "Its a description testing",
	Price: 20.50,
	ImageName: Image,
	Newsletter: 1,

}

func Test_AddProduct_MustReturnSuccess(t *testing.T) {

	ProductServiceMock := new(mocks.ProductRepository)

	ProductServiceMock.On("AddProduct", mock.Anything).Return(ProductData)


}
