package service

import (
	"context"

	checkout "gomall/app/frontend/hertz_gen/frontend/checkout"
	rpc "gomall/app/frontend/infra/rpc"
	frontendUtils "gomall/app/frontend/utils"
	rpccheckout "gomall/rpc_gen/kitex_gen/checkout"
	rpcpayment "gomall/rpc_gen/kitex_gen/payment"

	utils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId: uint32(userId),
		Email: req.Email,
		FirstName: req.Firstname,
		LastName: req.Lastname,
		Address: &rpccheckout.Address{
			StreetAddress: req.Street,
			City: req.City,
			State: req.Province,
			Country: req.Country,
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber: req.CardNum,
			CreditCardExpirationYear: req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
			CreditCardCvv: req.Cvv,
		},
	})
	if err != nil {
		klog.Error(err.Error())
		return nil, err
	}

	return utils.H{
		"title": "waiting",
		"redirect": "/checkout/result",
	}, nil
}