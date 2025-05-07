package main

import (
	"fmt"
	"os"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/model"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func main() {

	// 从环境变量获取MySQL配置
	dsn := fmt.Sprintf(
		conf.GetConfWithPath("../../conf").MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	fmt.Println(dsn)

	// 初始化连接
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	if err != nil {
		panic(err)
	}

	// CRUD
	DB.Create(&model.User{
		Email:    "demo@example.com",
		PassWord: "123",
	})

	DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("pass_word", "123456")

	var row model.User
	DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&row)
	fmt.Println(row)

	// 逻辑删除
	DB.Where("email = ?", "demo@example.com").Delete(&model.User{})

	// 物理删除
	DB.Unscoped().Where("email = ?", "demo@example.com").Delete(&model.User{})

}
