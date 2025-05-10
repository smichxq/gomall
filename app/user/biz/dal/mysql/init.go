package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/gomall/app/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// 从环境变量获取MySQL配置
	// 将环境变量注入到配置文件
	dsn := fmt.Sprintf(
		conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	fmt.Println(dsn)

	// 加载MySQL配置
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
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

	// 自动迁移
	// DB.AutoMigrate(&model.User{})

	fmt.Println(v)
}

// 适配单元测试中获取目录的问题
func InitUnitTest() {
	// 从环境变量获取MySQL配置
	// 将环境变量注入到配置文件
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	fmt.Println(dsn)

	// 加载MySQL配置
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
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

	// 自动迁移
	// DB.AutoMigrate(&model.User{})

	fmt.Println(v)
}
