package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"type:varchar(255);uniqueIndex"`
	PasswordHashed string `grom:"type:varchar(255);not null"`
}

// 设置表名
func (User) TableName() string {
	return "user"
}

// CRUD
func Create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
