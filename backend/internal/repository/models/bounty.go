package models

import "bounty_board/internal/constants"

type Bounty struct {
	ID          int                    `gorm:"primaryKey"`
	RoomID      int                    `gorm:"not null;index"`
	AuthorID    int                    `gorm:"not null;index"`
	ExecutorID  *int                   `gorm:"index"`
	Title       string                 `gorm:"size:128;not null"`
	Description string                 `gorm:"type:text;not null"`
	Reward      int                    `gorm:"not null"`
	Status      constants.BountyStatus `gorm:"size:32;not null;default:open"`

	Room        Room         `gorm:"foreignKey:RoomID"`
	Author      User         `gorm:"foreignKey:AuthorID"`
	Executor    *User        `gorm:"foreignKey:ExecutorID"`
	Submissions []Submission `gorm:"foreignKey:BountyID"`
}
