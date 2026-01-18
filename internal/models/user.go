package models

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;size:50" json:"id"`
	Username  string    `gorm:"uniqueIndex;not null;size:50" json:"username"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Todos     []Todo    `gorm:"foreignKey:UserID;references:ID" json:"todos,omitempty"`
}

func (User) TableName() string {
	return "users"
}
