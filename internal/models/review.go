package models

import "time"

type Review struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	GameID    uint      `json:"game_id"`
	UserID    uint      `json:"user_id"`
	Content   string    `json:"content"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}
