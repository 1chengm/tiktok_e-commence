package model

import (
	"context"
	"gorm.io/gorm"
)
type Category struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`

	Products []Product `gorm:"many2many:product_category;" json:"products"`
}

func (c Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db *gorm.DB
}
func (c CategoryQuery) GetProductByCategoryName(name string) (categories []Category, err error) {

	err = c.db.WithContext(c.ctx).Model(&Category{Name: name}).Where(&Category{Name: name}).Preload("Products").Find(&categories).Error
	return 
}
func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db: db,
	}
}