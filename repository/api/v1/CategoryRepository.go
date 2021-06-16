package v1

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vipulsirdemo/models"
	"vipulsirdemo/redis"
	v1res "vipulsirdemo/resources/response/api/v1"
)

type CategoryRepository interface {
	AddCategory(category models.Category)
	EditCategory(category models.Category)
}

type CategoryRepo struct {
	DB *gorm.DB
}

func (cr *CategoryRepo) GetCategoryById(Id uint) (models.Category, error) {
	category := models.Category{}
	err := cr.DB.Model(category).Select("id, name, status").Where("id = ?", Id).First(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (cr *CategoryRepo) AddCategory(category models.Category) {
	cr.DB.Model(&category).Create(&category)
	cr.AddCategoryToRedis()
	return
}

func (cr *CategoryRepo) EditCategory(category models.Category) {
	cr.DB.Model(&category).Update(&category)
	cr.AddCategoryToRedis()
	return
}

func (cr *CategoryRepo) GetAllCategories() ([]v1res.CategoryResponse, error) {
	var categoryList []v1res.CategoryResponse
	cr.DB.Table("categories").Select("categories.*").Where("status = ?", "active").Find(&categoryList)
	for ind, res := range categoryList {
		var menuList []v1res.CategoryMenuResponse
		cr.DB.Table("menus").Select("*").Where("category_id = ?", res.Id).Where("status = ?", "active").Find(&menuList)
		categoryList[ind].Menu = menuList
	}
	return categoryList, nil
}

func (cr *CategoryRepo) AddCategoryToRedis() {
	categoryList, err := cr.GetAllCategories()
	if err != nil {
		fmt.Println(err)
	}
	redis.SetValue("categories", categoryList)

}

func (cr *CategoryRepo) GetCategoryFromRedis() string {
	res := redis.GetValue("categories")
	fmt.Println("string categories: ", res)
	fmt.Printf("type %T: ", res)
	return "res"
}
