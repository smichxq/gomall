package utils

import (
	"context"

	frontendUtils "github.com/cloudwego/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

// 提取user_id
func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	// 从中间件中获取userId
	content["user_id"] = frontendUtils.GetUserIdFromCtx(ctx)

	return content
}
