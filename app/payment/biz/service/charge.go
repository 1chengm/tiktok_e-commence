package service

import (
	"context"
	"gomall/app/payment/biz/dal/mysql"
	payment "gomall/rpc_gen/kitex_gen/payment"
	creditcard "github.com/durango/go-credit-card"
	"gomall/app/payment/biz/model"
	"strconv"
	"time"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	card := creditcard.Card{
		Number:  req.CreditCard.CreditCardNumber,
		Cvv:	 strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:	 strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}
	err = card.Validate(true)
	if err != nil {
		return nil, kerrors.NewBizStatusError(4004001,err.Error())
	}
	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewBizStatusError(4005001,err.Error())
	}
	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId: 		req.UserId,
		OrderId:    	req.OrderId,
		TransActionId: 	transactionId.String(),
		Amount: 		req.Amount,
		PayAt: 			time.Now(),
	})
	if err != nil {
		return nil, kerrors.NewBizStatusError(4005002,err.Error())
	}
	return &payment.ChargeResp{ TransactionId: transactionId.String()}, nil
}
