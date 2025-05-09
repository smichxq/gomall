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

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从session获取认证信息
		s := sessions.Default(c)

		ctx = context.WithValue(ctx, SessionUserId, s.Get("user_id"))

		c.Next(ctx)
	}
}
