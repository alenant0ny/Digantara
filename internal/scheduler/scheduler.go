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

func AddJob(name, cronExpr, message string) (cron.EntryID, error) {
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
		JobName:   "EmailJob",
		JobType:   "",
		CronExpr:  "0 0 * * *",      // every day at midnight
		LastRun:   time.Time{},      // zero value if not run yet
		NextRun:   c.Entry(id).Next, // example next run
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

func ListJobs() {
	for name, id := range jobIDs {
		entry := c.Entry(id)
		fmt.Printf("Job: %s → Next run: %v\n", name, entry.Next)
	}
}

func JobDetails() {
	fmt.Println("job details by id")
}

func StartDbJobs() {
	fmt.Println("starting jobs from db after application restart")
}
