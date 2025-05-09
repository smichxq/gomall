package dal

import (
	"github.com/cloudwego/gomall/app/user/biz/dal/mysql"
	"github.com/cloudwego/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
