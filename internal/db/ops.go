package db

import (
	"gorm.io/gorm"
)

type JobRepository struct {
	DB *gorm.DB
}

// CreateJob inserts a new job record into the DB
func CreateJob(job *Job) error {
	return DB.Create(job).Error
}
