package service

import (
	"context"
	"strconv"

	common "gomall/app/frontend/hertz_gen/frontend/common"
	"gomall/app/frontend/infra/rpc"
	frontendUtils "gomall/app/frontend/utils"
	rpccart "gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "gomall/rpc_gen/kitex_gen/product"
	utils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// show the checkout page, and the user can choose the payment method
	var items []map[string]string
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: uint32(userId),
	})
	var total float32
	if err != nil {
		return nil, err
	}
	for _, item := range carts.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name": p.Name,
			"Price": strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture": p.Picture,
			"Qty": strconv.Itoa(int(item.Quantity)),
		})
		total += float32(item.Quantity) * p.Price
	}
	return utils.H{
		"title": "Checkout",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	} , nil
}
