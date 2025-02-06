package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "T-shirt-1", "Price": "100$","Picture":"/static/image/t-shirt-1.jpeg"},
		{"Name": "T-shirt-2", "Price": "110$","Picture":"/static/image/t-shirt-1.jpeg"},
		{"Name": "T-shirt-3", "Price": "120$","Picture":"/static/image/t-shirt-2.jpeg"},
		{"Name": "T-shirt-4", "Price": "130$","Picture":"/static/image/t-shirt-2.jpeg"},
		{"Name": "sweateshirt", "Price": "140$","Picture":"/static/image/sweatshirt.jpeg"},
		{"Name": "notebook", "Price": "150$","Picture":"/static/image/notebook.jpeg"},
		
	}
	resp["Title"] = "Home Page!"
	resp["Items"] = items
	
	return resp, nil
}
