package service

import (
	"context"
	"gomall/app/order/biz/dal/mysql"
	"gomall/app/order/biz/model"
	"gomall/rpc_gen/kitex_gen/cart"
	order "gomall/rpc_gen/kitex_gen/order"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	list, err := model.ListOrder(s.ctx,mysql.DB,req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500001, err.Error())
	}
	var orders []*order.Order
	for _, v := range list {
		var items []*order.OrderItem
		for _, vi := range v.OrderItems{
			items = append(items, &order.OrderItem{
				Item: &cart.GetItem{
					ProductId: vi.ProductId,
					Quantity: vi.Quantity,
				},
				Cost: vi.Cost,
			})
		}
	
		orders = append(orders, &order.Order{
			OrderId: v.OrderId,
			UserId: uint32(v.UserId),
			UserCurrency: v.UserCurrency,
			CreatedAt:    int32(v.CreatedAt.Unix()),
			Address: &order.Address{
				StreetAddress: v.Consignee.StreetAddress,
				City: v.Consignee.City,
				State: v.Consignee.State,
				Country: v.Consignee.Country,
				ZipCode: uint32(v.Consignee.ZipCode),
			},
			Items: items,
			Email: v.Consignee.Email,
		})
		resp = &order.ListOrderResp{
			Orders: orders,
		}
	}
	return
}