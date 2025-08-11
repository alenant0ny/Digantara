package scheduler

import (
	"digantara/internal/db"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

var c = cron.New(cron.WithSeconds())
var jobIDs = map[string]cron.EntryID{}

func StartScheduler() {
	c.Start()
}

func AddJob(name, cronExpr, message, jobname string) (cron.EntryID, error) {
	job, err := GetJob(name)
	if err != nil {
		return 0, err
	}
	id, err := c.AddFunc(cronExpr, func() { job.Run(message) })
	if err != nil {
		return 0, err
	}

	newJob := db.Job{
		JobID:     int(id),
		JobName:   jobname,
		JobType:   name,
		CronExpr:  cronExpr,
		LastRun:   c.Entry(id).Prev,
		NextRun:   c.Entry(id).Next,
		CreatedAt: time.Now(),
		Message:   message,
	}

	dbErr := db.CreateJob(&newJob)
	if dbErr != nil {
		fmt.Println(dbErr)
		return 0, dbErr
	}
	jobIDs[name] = id
	return jobIDs[name], nil
}

func JobDetails() {
	fmt.Println("job details by id")
}

func StartDbJobs() {
	fmt.Println("starting jobs from db after application restart")
}
