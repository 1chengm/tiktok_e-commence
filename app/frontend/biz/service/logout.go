package service

import (
	"context"

	"github.com/hertz-contrib/sessions"
	"github.com/cloudwego/hertz/pkg/app"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *common.Empty) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo logout-svr
	session := sessions.Default(h.RequestContext)
	
	session.Clear()
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
