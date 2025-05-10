package service

import (
	"context"
	"os"
	"reflect"
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

func TestRegisterService_Run(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		req *user.RegisterReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *user.RegisterResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RegisterService{
				ctx: tt.fields.ctx,
			}
			gotResp, err := s.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("RegisterService.Run() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
