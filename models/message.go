package models

import "time"

type Message struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   // foreign key to User.ID
	Username  string // denormalized for convenience
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
}
