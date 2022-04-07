package model

import "time"

type User struct {
	ID        uint      `gorm:"primary_key"`
	UUID      string    `json:"-"`
	IsValid   uint      `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
