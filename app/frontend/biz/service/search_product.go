package service

import (
	"context"
	rpcproduct "gomall/rpc_gen/kitex_gen/product"
	product "gomall/app/frontend/hertz_gen/frontend/product"
	"gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductService(Context context.Context, RequestContext *app.RequestContext) *SearchProductService {
	return &SearchProductService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductService) Run(req *product.SearchReq) (resp map[string]any, err error) {
	// todo edit your code
	p,err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{
		Query: req.Q,
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"items": p.Results,
		"q": req.Q,
	}, nil
}
