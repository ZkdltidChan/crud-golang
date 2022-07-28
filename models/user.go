package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID        string         `gorm:"primaryKey" json:"id"`
	Username  string         `json:"username" binding:"required"`
	Password  string         `json:"password" binding:"required"`
	Email     string         `json:"email" binding:"required"`
	NickName  string         `json:"nick_name"`
	CreatedAt int64          `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (b *User) TableName() string {
	return "user"
}

func (x *User) FillDefaults() {
	if x.ID == "" {
		x.ID = uuid.New().String()
	}
}
