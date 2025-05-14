package service

import (
	"context"

	auth "gomall/app/frontend/hertz_gen/frontend/auth"
	"gomall/app/frontend/infra/rpc"
	"gomall/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo login-svr
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return "", err
	}

	session := sessions.Default(h.RequestContext)

	session.Set("user_id", resp.UserId)
	session.Set("jwt_token", resp.Token) // Store the JWT token
	err = session.Save()
	if err != nil {
		hlog.Error("Failed to save session", err)
		return "", err
	}
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}

	return redirect, nil
}
