package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type PaymentLog struct {
	gorm.Model
	UserId uint32
	OrderId string
	TransActionId string // 交易ID
	Amount float32 // 金额
	PayAt time.Time // 支付时间
}


func TableName(p PaymentLog) string {
	return "payment_log"
}

func CreatePaymentLog(db *gorm.DB,ctx context.Context, payment *PaymentLog) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(payment).Error
}