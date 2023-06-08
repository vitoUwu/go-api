package models

import "time"

type UserModel struct {
	ID        uint      `gorm:"primaryKey;->;not null" json:"id"`
	Username  string    `gorm:"unique" json:"username"`
	CreatedAt time.Time `gorm:"autoCreateTime;->" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
