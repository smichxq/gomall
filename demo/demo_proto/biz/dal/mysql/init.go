package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {

	// 从环境变量获取MySQL配置
	dsn := fmt.Sprintf(
		conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	// 从配置文件加载MySQL配置
	// DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
	// 	&gorm.Config{
	// 		PrepareStmt:            true,
	// 		SkipDefaultTransaction: true,
	// 	},
	// )

	if err != nil {
		panic(err)
	}

	type Version struct {
		Version string
	}

	var v Version

	err = DB.Raw("select version() as version").Scan(&v).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(v)

}
