package models

import "time"

type FavoriteBook struct {
	UserID    uint `json:"user_id" gorm:"primaryKey"`
	BookID    uint `json:"book_id" gorm:"primaryKey"`
	CreatedAt time.Time
}
