package scheduler

import (
	"digantara/internal/db"
	"fmt"
	"log"
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

func StartDbJobs() {
	jobs, err := db.GetAllJobs()
	if err != nil {
		log.Fatalf("Unable to run jobs from db")
	}
	for _, v := range jobs {
		job, err := GetJob(v.JobType)
		if err != nil {
			log.Fatalf("Could not get jobs")
		}
		go func(job Job, v db.Job) {
			_, err := c.AddFunc(v.CronExpr, func() { job.Run(v.Message) })
			if err != nil {
				log.Fatalf("Could not start job of id: %v", v.ID)
			}
		}(job, v)

	}
	fmt.Println(">>>>>>>>>>>>>>>", jobs)
	fmt.Println("starting jobs from db after application restart")
}
