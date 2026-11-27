package models

type User struct {
	ID           int    `gorm:"primaryKey"`
	Username     string `gorm:"size:64;not null;uniqueIndex"`
	Email        string `gorm:"size:255;not null;uniqueIndex"`
	PasswordHash string `gorm:"size:255;not null"`
}
