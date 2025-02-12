package service

import (
	"context"
	"gomall/app/cart/biz/model"
	"gomall/app/cart/rpc"
	"gomall/app/cart/biz/dal/mysql"
	cart "gomall/rpc_gen/kitex_gen/cart"
	"gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	productRsep ,err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}
	if productRsep == nil || productRsep.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(404, "product not found")
	}
	cartItem := &model.Cart{
		UserID:    req.UserId,
		ProductID: req.Item.ProductId,
		Qty:       req.Item.Quantity,
	}
	err = model.AddItem(s.ctx, mysql.DB, cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}
	return &cart.AddItemResp{}, nil
}
