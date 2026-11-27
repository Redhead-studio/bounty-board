package models

import (
	"bounty_board/internal/constants"
	"time"
)

type RoomMember struct {
	ID       int                        `gorm:"primaryKey"`
	RoomID   int                        `gorm:"not null;uniqueIndex:idx_room_member"`
	UserID   int                        `gorm:"not null;uniqueIndex:idx_room_member"`
	Role     constants.RoomMemberRole   `gorm:"size:32;not null;default:member"`
	Status   constants.RoomMemberStatus `gorm:"size:32;not null;default:active"`
	JoinedAt time.Time                  `gorm:"not null;autoCreateTime"`

	Room Room `gorm:"foreignKey:RoomID"`
	User User `gorm:"foreignKey:UserID"`
}
