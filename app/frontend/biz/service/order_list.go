package service

import (
	"context"

	common "gomall/app/frontend/hertz_gen/frontend/common"
	rpc "gomall/app/frontend/infra/rpc"
	"gomall/app/frontend/types"
	frontendutils "gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/order"
	"gomall/rpc_gen/kitex_gen/product"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	//"github.com/cloudwego/hertz/pkg/common/hlog"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{UserId: uint32(userId)})
	//实际中应该分步请求，不应该一次性请求所有数据,
	if err != nil {
		return nil, err
	}
	var list []types.Order
	for _ , v := range orderResp.Orders {
		var(
			 total float32
			 items 	[]types.OrderItem
		)
		for _, vv := range v.Items {
			total += vv.Cost

			i := vv.Item
			productResp , err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: i.ProductId})
			if err != nil {
				return nil, err
			}
			if productResp.Product == nil || productResp == nil {
				continue
			}
			items = append(items, types.OrderItem{
				Picture: productResp.Product.Picture,
				ProductName: productResp.Product.Name,
				Qty: i.Quantity,
				Cost: vv.Cost,
			})
			
		}	
		createdAt := time.Unix(int64(v.CreatedAt), 0)
		list = append(list, types.Order{
			OrderId: v.OrderId,
			CreatedDate: createdAt.Format("2006-01-02 15:04:05"),
			Cost:  total,
			Items: items,
		})
			
		}
		return utils.H{
			"title": "Order List",
			"orders": list,
		}, nil
}

