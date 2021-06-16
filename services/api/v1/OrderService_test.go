package v1Service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"vipulsirdemo/mocks"
	"vipulsirdemo/models"
	v1req "vipulsirdemo/resources/request/api/v1"
	v1res "vipulsirdemo/resources/response/api/v1"
)

var (
	id uint = 1
)


var OrderRequest = v1req.OrderRequest{
	ID:          1,
	UserId:      1,
	OrderNumber: "001",
	Description: "Testing Description",
	Date:        time.Time{},
	Status:      "active",
}
var OrderData = models.Order{
	Model: models.Model{
		ID: 1,
	},
	UserId:      1,
	OrderNumber: "123",
	Description: "Testing description",
	Status:      "active",
}
var OrderResponse = v1res.OrderResponse{
	OrderNumber: "001110",
	Description: "testing description",
	Status:      "active",
	UserId:      1,
}


func Test_AddOrders_MustReturnSuccess(t *testing.T) {
	OrderData := map[string]interface{}{
		"data": v1req.OrderRequest{
			ID:   1,
			UserId:  1,
			OrderNumber: "1234567890",
			Description:  "This is the testing",
			Date: time.Now(),
			Status: "active",
		},
		"message": "Order added successfully.",
		"status":  1,
	}
	AddOrderMock := new(mocks.OrderRepository)

	AddOrderMock.On("AddOrder", mock.Anything).Return(OrderData)

	var srvc OrderService
	err := srvc.AddOrder(OrderRequest)
	if err != nil {
		assert.Equal(t, err, OrderData)
	}
}

/*func Test_GetOrders_MustReturnSuccess(t *testing.T) {
	TestGetOrderMock := new(mocks.OrderRepository)

	TestGetOrderMock.On("GetOrderByUser", mock.Anything).Return(id)

}
*/