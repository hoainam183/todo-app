package model

import "time"

type Todo struct {
	ID          string    `gorm:"primaryKey;size:50" json:"id"`
	UserID      string    `gorm:"not null;index" json:"user_id"`
	Title       string    `gorm:"not null;size:255" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Completed   bool      `gorm:"default:false" json:"completed"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Foreign key relationship
	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
}

// TableName specifies the table name for Todo model
func (Todo) TableName() string {
	return "todos"
}
