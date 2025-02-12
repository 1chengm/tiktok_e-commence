package service

import (
	"context"
	"gomall/app/cart/biz/dal/mysql"
	"gomall/app/cart/biz/model"
	cart "gomall/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
}

// NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	list, err := model.GetCartByUserID(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, err
	}
	var items []*cart.GetItem
	for _, v := range list {
		items = append(items, &cart.GetItem{
			ProductId: v.ProductID,
			Quantity:  v.Qty,
		})
	}
	return &cart.GetCartResp{Items: items}, nil
}
