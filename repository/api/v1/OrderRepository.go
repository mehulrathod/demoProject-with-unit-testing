package v1

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vipulsirdemo/models"
	v1res "vipulsirdemo/resources/response/api/v1"
)

type OrderRepository interface {
	AddOrder(order models.Order)
	GetOrder(order models.Order)
}

type OrderRepo struct {
	DB *gorm.DB
}

func (or *OrderRepo) AddOrder(order models.Order) {
	or.DB.Model(&order).Create(&order)
}

func (or *OrderRepo) GetOrderByUser(id uint) ([]v1res.OrderResponse, error) {
	var orderList []v1res.OrderResponse
	or.DB.Table("orders").Select("orders.*").Where("user_id = ? ", id).Find(&orderList)
	for ind, res := range orderList {
		var orderMenuList []v1res.OrderMenuResponse
		or.DB.Table("order_menu").Select("order_menu.*").Where("order_id = ? ", res.Id).Find(&orderMenuList)
		for i, r := range orderMenuList {
			var menuList v1res.CategoryMenuResponse
			or.DB.Table("menus").Select("*").Where("id = ?", r.MenuId).Where("status = ?", "active").Find(&menuList)
			orderMenuList[i].Menu = menuList
			orderMenuList[i].Price = menuList.Price
		}

		orderList[ind].Items = orderMenuList

		fmt.Println("latest", orderMenuList)
		fmt.Println("result", res.Id)
	}
	return orderList, nil
}
