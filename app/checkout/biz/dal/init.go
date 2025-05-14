package dal

import (
	"github.com/cloudwego/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	// mysql.Init()
}
