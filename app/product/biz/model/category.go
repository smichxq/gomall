package model

import (
	"context"

	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`

	Products []Product `json:"product" gorm:"many2many:product_category"`
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// 迁移时指定表名
func (c Category) TableName() string {
	return "category"
}

func (c CategoryQuery) GetProductsByCategoryName(name string) (categories []Category, err error) {
	// Preload自动关联Products中符合Category查询的结果
	err = c.db.WithContext(c.ctx).Model(&Category{}).Where("nema = ?", name).Preload("Products").Find(&categories).Error

	return
}
