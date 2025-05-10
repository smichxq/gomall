package service

import (
	"context"
	"os"
	"testing"

	"github.com/cloudwego/gomall/app/user/biz/dal/mysql"
	user "github.com/cloudwego/gomall/rpc_gen/kitex_gen/user"
)

func TestRegister_Run(t *testing.T) {
	// 临时测试环境变量
	os.Setenv("MYSQL_USER", "root")
	os.Setenv("MYSQL_PASSWORD", "123")
	os.Setenv("MYSQL_HOST", "192.168.3.6")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_DATABASE", "user")

	mysql.InitUnitTest()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "example@aaa.com",
		Password:        "123asd",
		PasswordConfirm: "123asd",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
