package model

// OrderItem 订单商品关联表
import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ProductId uint32 `gorm:"type:int(11)"`
	OrderIdRefer string `gorm:"type:varchar(100);index;not null"`
	Quantity uint32 `gorm:"type:int(11);not null"`
	Cost float32 `gorm:"type:decimal(10,2);not null"`
}
func (OrderItem)TableName() string {
	return "order_item"
}