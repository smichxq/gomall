package middleware

import "github.com/cloudwego/hertz/pkg/app/server"

// 中间件注册器
func Register(h *server.Hertz) {
	h.Use(GlobalAuth())
}
