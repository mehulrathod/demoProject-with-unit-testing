package v1

import (
	"github.com/jinzhu/gorm"
	"vipulsirdemo/models"
	v1res "vipulsirdemo/resources/response/api/v1"
)

type MenuRepository interface {
	AddMenu(menu models.Menu)
	EditMenu(menu models.Menu)
}

type MenuRepo struct {
	DB *gorm.DB
}

func (mr *MenuRepo) GetMenuById(Id uint) (models.Menu, error) {
	menu := models.Menu{}
	err := mr.DB.Model(menu).Select("id, name, status, price, image, category_id").Where("id = ?", Id).First(&menu).Error
	if err != nil {
		return menu, err
	}
	return menu, nil
}

func (mr *MenuRepo) GetMenuByName(name string, id uint) (models.Menu, error) {
	menu := models.Menu{}
	err := mr.DB.Model(menu).Select("id, name, description, status, price, image, category_id").Where("category_id = ?", id).Where("name = ?", name).First(&menu).Error
	if err != nil {
		return menu, err
	}
	return menu, nil
}

func (mr *MenuRepo) AddMenu(menu models.Menu) {
	mr.DB.Model(&menu).Create(&menu)
	return
}

func (mr *MenuRepo) EditMenu(menu models.Menu) {
	mr.DB.Model(&menu).Update(&menu)
	return
}
func (mr *MenuRepo) GetAllMenu() ([]v1res.MenuResponse, error) {
	var menuList []v1res.MenuResponse
	mr.DB.Table("menus").Select("menus.*, categories.name as category").Joins("JOIN categories on categories.id = menus.category_id").Find(&menuList)
	return menuList, nil
}
