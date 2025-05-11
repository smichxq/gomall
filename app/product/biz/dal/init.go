package dal

import (
	"github.com/cloudwego/gomall/app/product/biz/dal/mysql"
	"github.com/cloudwego/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
