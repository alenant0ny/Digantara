package db

import (
	"time"
)

type Job struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	JobID     int    `gorm:"not null"` // your custom sequential job id
	JobName   string `gorm:"not null"`
	JobType   string
	CronExpr  string `gorm:"not null"`
	LastRun   time.Time
	NextRun   time.Time
	CreatedAt time.Time
	Message   string
}
