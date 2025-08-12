package db

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type JobRepository struct {
	DB *gorm.DB
}

// CreateJob inserts a new job record into the DB
func CreateJob(job *Job) error {
	return DB.Create(job).Error
}

func GetAllJobs() ([]Job, error) {
	var jobs []Job
	result := DB.Find(&jobs)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return jobs, nil
}

func GetJobByID(id int) (Job, error) {
	var job Job
	result := DB.Where("job_id = ?", id).First(&job)
	if result.Error != nil {
		return job, errors.New("Job not found")
	}

	return job, nil
}

func UpdateJobByID(job *Job) {

	result := DB.Model(&Job{}).
		Where("job_id = ?", job.ID).
		Updates(Job{
			LastRun: job.LastRun,
			NextRun: job.NextRun,
			JobID:   job.JobID,
		})
	log.Println(result)

}
