package dal

import (
	"github.com/cloudwego/gomall/app/notify/biz/dal/mysql"
	"github.com/cloudwego/gomall/app/notify/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
