package model

import (
	"context"
	"gorm.io/gorm"
)
type Product struct {
	Base
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	Picture string `json:"picture"`

	Categories []Category `gorm:"many2many:product_category;" json:"categories"`
}
func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx   context.Context
	db   *gorm.DB
}
func (p *ProductQuery) GetByID(id uint) (product Product, err error) {

	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, id).Error
	return 
}
func (p *ProductQuery) SearchProducts(name string) (products []Product, err error) {

	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products,"name Like ? or description Like ?",
		"%"+name+"%","%"+name+"%").Error
	return 
}
func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db: db,
	}
}