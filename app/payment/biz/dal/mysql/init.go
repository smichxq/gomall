package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/gomall/app/payment/biz/model"
	"github.com/cloudwego/gomall/app/payment/conf"

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
	migrate()

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
	migrate()

	fmt.Println(v)
}

func migrate() {
	// panic("impl")
	if os.Getenv("GO_ENV") != "online" {
		// 检查数据库中是否已存在表
		needDemoData := !DB.Migrator().HasTable(&model.PaymentLog{})

		// 自动迁移数据库表结构
		// 根据模型定义创建或更新表结构
		// 如果存在则尝试添加缺失字段、索引等，但不会删除字段或索引
		// 多个模型可以一次性迁移
		// 也会创建关联表
		err := DB.AutoMigrate(
			&model.PaymentLog{},
		)
		if err != nil {
			panic("migrate fail")
		}

		if needDemoData {
		}

	}
}
