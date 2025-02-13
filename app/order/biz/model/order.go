package model

import( 
	"gorm.io/gorm"
	"context"
)

//收货人信息
type Consignee struct {
	Email string `json:"email"`
	StreetAddress string `json:"street"`
	City string 
	State string
	Country string
	ZipCode int32

}


type Order struct {
	gorm.Model
	OrderId string `gorm:"type:varchar(100);uniqueIndex;not null"`
	UserId uint `gorm:"type:varchar(100);index;not null"`
	UserCurrency string `gorm:"type:varchar(10)"`
	Consignee Consignee `gorm:"embedded"`//收货人信息,嵌入结构体
	OrderItems []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`//订单商品关联表
}


func (Order)TableName() string {
	return "order"
}
func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ( []*Order, error) {
	// Finish your business logic.
	 var orders []*Order
	 err := db.WithContext(ctx).Where("user_id = ?", userId).Preload("OrderItems").Find(&orders).Error//关联查询
	 if err != nil {
		return nil, err
	 }
	return orders, nil
}