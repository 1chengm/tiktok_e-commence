package middleware

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(h *server.Hertz) {
	// todo edit custom code
	h.Use(GlobalAuth())
}