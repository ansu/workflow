package model

import (
	"time"
)

type Workflow struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
