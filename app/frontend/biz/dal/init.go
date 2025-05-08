package dal

import (
	"github.com/cloudwego/gomall/app/frontend/biz/dal/mysql"
	"github.com/cloudwego/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
