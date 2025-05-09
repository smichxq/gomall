package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex"`
	PasswordHashed string `grom:"type:varchar(255) not null"`
}

// 设置表名
func (User) TableName() string {
	return "user"
}
