package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudwego/gomall/app/user/biz/dal/mysql"
	"github.com/cloudwego/gomall/app/user/biz/model"
	user "github.com/cloudwego/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.

	// 获取请求的用户名
	userEmail, userPasswd := req.Email, req.Password

	fmt.Println(userEmail, userPasswd)

	if userEmail == "" || userPasswd == "" {
		return nil, errors.New("email or passwd is empty")
	}

	// 查询记录
	user_row, err := model.SelectByEmail(mysql.DB, userEmail)
	if err != nil {
		return nil, err
	}

	// 加密
	passwdHashed, err := bcrypt.GenerateFromPassword([]byte(userPasswd), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 比对
	if string(passwdHashed) != user_row.PasswordHashed {
		return nil, err
	}

	return &user.LoginResp{UserId: int32(user_row.ID)}, nil
}
