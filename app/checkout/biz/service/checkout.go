package service

import (
	"context"
	"gomall/app/checkout/infra/rpc"
	cart "gomall/rpc_gen/kitex_gen/cart"
	checkout "gomall/rpc_gen/kitex_gen/checkout"
	payment "gomall/rpc_gen/kitex_gen/payment"
	product "gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/google/uuid"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Items == nil || len(cartResult.Items) == 0 {
		return nil, kerrors.NewBizStatusError(5004001, "cart is empty")
	}
	var total float32
	for _, item := range cartResult.Items {
		productResp, resulterr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id : item.ProductId,
		})
		if resulterr != nil {
			return nil, kerrors.NewBizStatusError(5005001, resulterr.Error())
		}
		if productResp.Product == nil {
			return nil, kerrors.NewBizStatusError(5004001, "product not found")
		}
		 p := productResp.Product
		cost := p.Price * float32(item.Quantity)
		total += cost
	}
	var  orderId string
	u, _ := uuid.NewRandom()
	orderId = u.String()
	payReq := &payment.ChargeReq{
		UserId: req.UserId,
		OrderId: orderId,
		Amount: total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber: req.CreditCard.CreditCardNumber,
			CreditCardCvv: req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear: req.CreditCard.CreditCardExpirationYear,
		},
	}
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		 klog.Error(err.Error())
	}
	paymentresult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, kerrors.NewBizStatusError(5005001, err.Error())
	}
	klog.Info("paymentResp:", paymentresult)
	resp = &checkout.CheckoutResp{
		OrderId: orderId,
		TransactionId: paymentresult.TransactionId,
	}
	return
}
