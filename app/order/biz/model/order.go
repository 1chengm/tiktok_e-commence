package model

import "gorm.io/gorm"

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
	Currency string `gorm:"type:varchar(10)"`
	Consignee string `gorm:"embedded"`//收货人信息,嵌入结构体
	OrderItems []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`//订单商品关联表
}


func (Order)TableName() string {
	return "order"
}