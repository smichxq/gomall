package service

import (
	"context"
	"log"

	auth "github.com/cloudwego/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
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
	// todo edit your code

	// 使用rpc调用服务端(app/user)
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}

	session := sessions.Default(h.RequestContext)

	session.Set("user_id", resp.UserId)

	err = session.Save()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// 如果存在Next则返回重定向的路径
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}

	return redirect, nil
}
