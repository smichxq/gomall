package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"type:varchar(255);uniqueIndex"`
	PasswordHashed string `grom:"type:varchar(255);not null"`
}

// 迁移时指定表名
func (User) TableName() string {
	return "user"
}

// CRUD
func Create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// CRUD
func SelectByEmail(db *gorm.DB, email string) (*User, error) {
	var user User

	fmt.Println(email)

	err := db.Where("email = ?", email).First(&user).Error

	return &user, err
}
