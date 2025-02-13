package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"

	frontendUtils "gomall/app/frontend/utils"
	
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// todo edit custom code
		s := sessions.Default(c)
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, s.Get("user_id"))
		c.Next(ctx)
	}
}
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// todo edit custom code
		s := sessions.Default(c)
		if s.Get("user_id") == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			//c.String(401, "please login")
			return
		}
		// ctx = context.WithValue(ctx, SessionUserId, s.Get("user_id"))
		c.Next(ctx)
	}
}