package v1Service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"vipulsirdemo/mocks"
	"vipulsirdemo/models"
	v1req "vipulsirdemo/resources/request/api/v1"
)

var (
	id uint = 1
)


func Test_AddOrders_MustReturnSuccess(t *testing.T)  {
	orderRepoMock := new(mocks.OrderRepository)

	var orderRequestParams v1req.OrderRequest

	orderRepoMock.On("AddOrder", orderRequestParams).Return(&models.Order{
		Model: models.Model{
			ID: 1,
		},
		UserId: 1,
		Status: "active",
		Description: "Description for testing",
		OrderNumber: "011",
	})

	var srvc OrderService
	err := srvc.AddOrder(orderRequestParams)
	assert.NotNil(t, err, nil)
}

func Test_AddOrders_Fail(t *testing.T)  {
	orderRepoMock := new(mocks.OrderRepository)

	var orderRequestParams v1req.OrderRequest

	orderRepoMock.On("AddOrder", orderRequestParams).Return(nil)

	var srvc OrderService
	err := srvc.AddOrder(orderRequestParams)
	assert.Nil(t, err, nil)
}