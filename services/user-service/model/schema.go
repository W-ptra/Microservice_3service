package model

import (
	"time"
)

type User struct {
	Id        int       	`gorm:"primaryKey;autoIncrement" json:"id"` 
	Name      string		`json:"name"`
	CreatedAt time.Time     `gorm:"autoCreateTime" json:"createdAt"` 
	UpdatedAt time.Time     `gorm:"autoUpdateTime" json:"updatedAt"`
}