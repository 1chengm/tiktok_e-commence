package service

import (
	"context"
	product "gomall/rpc_gen/kitex_gen/product"

	"gomall/app/product/biz/dal/mysql"
	"gomall/app/product/biz/model"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	p, err := productQuery.GetByID(uint(req.Id))
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(2004002, "product not found")
	}
	return &product.GetProductResp{
		Product: &product.Product{
		Id:          uint32(p.ID),
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Picture:     p.Picture,
		},
	}, err
}
