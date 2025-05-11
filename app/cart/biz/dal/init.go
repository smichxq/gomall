package dal

import (
	"github.com/cloudwego/gomall/app/cart/biz/dal/mysql"
	"github.com/cloudwego/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
