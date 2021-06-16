package v1

import (
	"github.com/jinzhu/gorm"
	"vipulsirdemo/models"
	v1res "vipulsirdemo/resources/response/api/v1"
)

type OrderMenuRepository interface {
	AddOrderMenu(orderMenu models.OrderMenu)
	GetOrderMenus(orderMenu models.OrderMenu)
}

type OrderMenuRepo struct {
	DB *gorm.DB
}

func (omr *OrderMenuRepo) AddOrderMenu(orderMenu models.OrderMenu) models.Menu {
	var menu models.Menu
	omr.DB.Table("menus").Select("menus.*").Where("id = ? ", orderMenu.MenuId).Find(&menu)
	orderMenu.Price = menu.Price
	omr.DB.Model(&orderMenu).Create(&orderMenu)
	return menu
}

func (omr *OrderMenuRepo) GetOrderMenus(id uint) ([]v1res.OrderMenuResponse, error) {
	var orderMenuList []v1res.OrderMenuResponse
	omr.DB.Table("order_menu").Select("order_menu.*").Where("order_id = ? ", id).Find(&orderMenuList)
	for ind, res := range orderMenuList {
		var menuList v1res.CategoryMenuResponse
		omr.DB.Table("menus").Select("*").Where("id = ?", res.MenuId).Where("status = ?", "active").Find(&menuList)
		orderMenuList[ind].Menu = menuList
		orderMenuList[ind].Price = menuList.Price
	}
	return orderMenuList, nil
}
