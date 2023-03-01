package model

import (
	"time"
)

type UserWorkflowScreen struct {
	ID         uint      `gorm:"primaryKey"`
	WorkflowID uint      `gorm:"not null"`
	ScreenID   uint      `gorm:"not null"`
	Status     string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`
}

// Status constants
const (
	StatusCreated   = "created"
	StatusAssigned  = "assigned"
	StatusCompleted = "completed"
	StatusCancelled = "cancelled"
)
