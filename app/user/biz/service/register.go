package service

import (
	"context"
	"errors"

	"github.com/cloudwego/gomall/app/user/biz/dal/mysql"
	"github.com/cloudwego/gomall/app/user/biz/model"
	user "github.com/cloudwego/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.

	// 校验
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password not match")
	}

	// 加密
	passwdHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 实例化
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwdHashed),
	}

	// 插入数据库
	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}

	// 返回id
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
