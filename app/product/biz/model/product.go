package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`

	Categories []Category `json:"categories" gorm:"many2many:product_category"`
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// 迁移时指定表名
func (p Product) TableName() string {
	return "product"
}

func Create(db *gorm.DB, product *Product) error {
	return db.Create(product).Error
}

func (query ProductQuery) QeryById(productId uint32) (product Product, err error) {
	// 使用context方便后续扩展(链路追踪)
	// Model指定查询表
	// First第二个参数通过主键查询
	err = query.db.WithContext(query.ctx).Model(&Product{}).First(&product, productId).Error
	return
}

func (query ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	// 将放q到两个占位符方便查询
	err = query.db.WithContext(query.ctx).Model(&Product{}).Find(&products, "name like ? or description like ?",

		"%"+q+"%", "%"+q+"%").Error

	return
}
