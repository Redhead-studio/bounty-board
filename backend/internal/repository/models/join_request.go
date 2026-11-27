package models

import (
	"bounty_board/internal/constants"
	"time"
)

type JoinRequest struct {
	ID         int                         `gorm:"primaryKey"`
	RoomID     int                         `gorm:"not null;index"`
	UserID     int                         `gorm:"not null;index"`
	Status     constants.JoinRequestStatus `gorm:"size:32;not null;default:pending"`
	CreatedAt  time.Time                   `gorm:"not null;autoCreateTime"`
	ReviewedAt *time.Time

	Room Room `gorm:"foreignKey:RoomID"`
	User User `gorm:"foreignKey:UserID"`
}
