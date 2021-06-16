package v1

import (
	"github.com/jinzhu/gorm"
	"vipulsirdemo/models"
)

type ProductRepository interface {
	AddProduct(product models.Products)
}

type ProductRepo struct {
	DB *gorm.DB
}

func (pr *ProductRepo) AddProduct(product models.Products) {
	pr.DB.Model(&product).Create(&product)
}
