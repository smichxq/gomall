package auth

import (
	"context"

	"github.com/cloudwego/gomall/app/frontend/biz/service"
	"github.com/cloudwego/gomall/app/frontend/biz/utils"
	auth "github.com/cloudwego/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Login .
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewLoginService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 登陆成功重定向

	c.Redirect(consts.StatusFound, []byte("/"))

	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
