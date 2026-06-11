package models

import (
	"time"

	"gorm.io/gorm"
)

type TestReport struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	TestCaseID uint      `json:"test_case_id"`
	TestCase   TestCase  `gorm:"foreignKey:TestCaseID" json:"test_case"`
	Status     string    `gorm:"default:'running'" json:"status"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	Response   string    `json:"response"`
	Error      string    `json:"error"`
	Duration   int64     `json:"duration"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}