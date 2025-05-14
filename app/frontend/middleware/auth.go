package middleware

import (
	"context"
	// "fmt" // For debugging, can be removed

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"
	frontendUtils "gomall/app/frontend/utils"
)

// GlobalAuth middleware: Populates user_id into context if a valid JWT token exists in session.
// This runs for all requests.
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		var userIDToSet interface{} // Store as interface{} initially

		tokenVal := s.Get("jwt_token")
		if tokenVal != nil {
			if tokenString, ok := tokenVal.(string); ok && tokenString != "" {
				claims, err := frontendUtils.ValidateJWT(tokenString) // Validate token
				if err == nil && claims != nil {
					userIDToSet = claims.UserID // Use UserID from valid token
					// Ensure session user_id is consistent with token
					currentSessionUserID := s.Get("user_id")
					if currentSessionUserID == nil || currentSessionUserID.(int32) != claims.UserID {
						s.Set("user_id", claims.UserID)
						_ = s.Save() // Best effort save, handle error if critical
					}
				} else {
					// Token is invalid or expired, clear it and user_id from session
					s.Delete("jwt_token")
					s.Delete("user_id")
					_ = s.Save()
				}
			}
		} else {
			// No JWT token, ensure user_id is also cleared if it exists without a token
			if s.Get("user_id") != nil {
				s.Delete("user_id")
				_ = s.Save()
			}
		}

		if userIDToSet != nil {
			// Ensure it's the correct type before setting in context
			if uid, ok := userIDToSet.(int32); ok {
				ctx = context.WithValue(ctx, frontendUtils.SessionUserId, uid)
			}
		}
		c.Next(ctx)
	}
}

// Auth middleware: Protects routes that require a valid logged-in user.
// Redirects to login if the token is missing or invalid.
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		hlog.CtxInfof(ctx, "GlobalAuth: session jwt_token: %v, session user_id: %v", s.Get("jwt_token"), s.Get("user_id"))
		tokenVal := s.Get("jwt_token")

		if tokenVal == nil {
			// fmt.Println("Auth: No token in session")
			c.Redirect(consts.StatusFound, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}

		tokenString, ok := tokenVal.(string)
		if !ok || tokenString == "" {
			// fmt.Println("Auth: Token in session is not a string or is empty")
			s.Delete("jwt_token") // Clean up bad token
			s.Delete("user_id")
			_ = s.Save()
			c.Redirect(consts.StatusFound, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}

		claims, err := frontendUtils.ValidateJWT(tokenString)
		if err != nil {
			// fmt.Printf("Auth: Token validation error: %v\n", err)
			s.Delete("jwt_token") // Clean up invalid/expired token
			s.Delete("user_id")
			_ = s.Save()
			c.Redirect(consts.StatusFound, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}

		// Token is valid. Ensure user_id in session matches the token's UserID.
		// GlobalAuth should have already populated context if token was valid.
		// This is an additional safeguard and ensures session consistency.
		currentSessionUserID := s.Get("user_id")
		if currentSessionUserID == nil || currentSessionUserID.(int32) != claims.UserID {
			s.Set("user_id", claims.UserID)
			_ = s.Save() // Best effort save
		}
		// If GlobalAuth hasn't run or to be absolutely sure context has user_id from validated token:
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, claims.UserID)

		c.Next(ctx)
	}
}
