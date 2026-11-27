package models

import "bounty_board/internal/constants"

type Submission struct {
	ID         int                        `gorm:"primaryKey"`
	BountyID   int                        `gorm:"not null;index"`
	ExecutorID int                        `gorm:"not null;index"`
	Result     string                     `gorm:"type:text;not null"`
	Status     constants.SubmissionStatus `gorm:"size:32;not null;default:pending"`

	Bounty   Bounty `gorm:"foreignKey:BountyID"`
	Executor User   `gorm:"foreignKey:ExecutorID"`
}
