package models

import (
	"time"

	"gorm.io/gorm"
)

type TestCase struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Protocol    string    `gorm:"default:'http'" json:"protocol"`
	Method      string    `gorm:"default:'GET'" json:"method"`
	URL         string    `gorm:"not null" json:"url"`
	Headers     string    `json:"headers"`
	Body        string    `json:"body"`
	Params      string    `json:"params"`
	Expected    string    `json:"expected"`
	Status      string    `gorm:"default:'draft'" json:"status"`
	CreatorID   uint      `json:"creator_id"`
	Creator     User      `gorm:"foreignKey:CreatorID" json:"creator"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}