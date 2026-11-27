package models

import "time"

type PointsTransaction struct {
	ID        int       `gorm:"primaryKey"`
	RoomID    int       `gorm:"not null;index"`
	UserID    int       `gorm:"not null;index"`
	BountyID  int       `gorm:"not null;uniqueIndex:idx_bounty_reward"`
	Amount    int       `gorm:"not null;check:chk_points_transactions_amount_positive,amount > 0"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`

	Room   Room   `gorm:"foreignKey:RoomID"`
	User   User   `gorm:"foreignKey:UserID"`
	Bounty Bounty `gorm:"foreignKey:BountyID"`
}
