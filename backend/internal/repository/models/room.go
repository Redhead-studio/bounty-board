package models

import "time"

type Room struct {
	ID         int       `gorm:"primaryKey"`
	Name       string    `gorm:"size:128;not null"`
	InviteCode string    `gorm:"size:64;not null;uniqueIndex"`
	MasterID   int       `gorm:"not null;index"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime"`

	Master       User          `gorm:"foreignKey:MasterID"`
	Members      []RoomMember  `gorm:"foreignKey:RoomID"`
	JoinRequests []JoinRequest `gorm:"foreignKey:RoomID"`
	Bounties     []Bounty      `gorm:"foreignKey:RoomID"`
}
