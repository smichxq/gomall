package model

import "time"

type Base struct {
	ID        int `gorm:"primary key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
