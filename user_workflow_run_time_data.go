package model

import "time"

type UserWorkflowRunTimeData struct {
	ID             uint   `gorm:"primaryKey"`
	ScreenID       uint   `gorm:"not null"`
	Name           string `gorm:"not null"`
	Type           string `gorm:"not null"`
	TextValue      string
	LongValue      int64
	DoubleValue    float64
	ByteArrayValue []byte
	DateValue      time.Time
	TextValue2     string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
