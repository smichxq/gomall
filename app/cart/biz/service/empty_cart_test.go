package service

import (
	"context"
	"os"
	"testing"

	"github.com/cloudwego/gomall/app/cart/biz/dal/mysql"
	"github.com/cloudwego/gomall/app/cart/infra/rpc"
	cart "github.com/cloudwego/gomall/rpc_gen/kitex_gen/cart"
)

func TestEmptyCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEmptyCartService(ctx)
	// init req and assert value
	// 临时测试环境变量
	os.Setenv("MYSQL_USER", "root")
	os.Setenv("MYSQL_PASSWORD", "123")
	os.Setenv("MYSQL_HOST", "192.168.3.6")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_DATABASE", "user")

	rpc.InitClientUnitTest("192.168.3.6:8500")
	mysql.InitUnitTest()

	req := &cart.EmptyCartReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
