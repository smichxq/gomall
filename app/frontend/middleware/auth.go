package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

// 定义session常量
type SessionUserIdKey string

// 同上
const SessionUserId SessionUserIdKey = "user_id"

// app.RequestContext获取session
// 写回context.Context
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从session获取认证信息
		s := sessions.Default(c)

		ctx = context.WithValue(ctx, SessionUserId, s.Get("user_id"))

		c.Next(ctx)
	}
}

// 鉴权特定路径
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)

		userId := s.Get("user_id")
		if userId == nil {
			c.Redirect(302, []byte("/sign-in"))
			// 终止后续中间件handler的调用
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
