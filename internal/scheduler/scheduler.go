package scheduler

import (
	"digantara/internal/db"
	"fmt"
	"log"
	"sync"
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

var wg sync.WaitGroup

func StartDbJobs() {
	fmt.Println("starting jobs from db after application restart")
	jobs, err := db.GetAllJobs()
	if err != nil {
		log.Fatalf("Unable to run jobs from db")
	}
	for _, v := range jobs {
		wg.Add(1)
		job, err := GetJob(v.JobType)
		if err != nil {
			log.Fatalf("Could not get jobs")
		}

		go func(job Job, v db.Job) {
			defer wg.Done()
			id, err := c.AddFunc(v.CronExpr, func() { job.Run(v.Message) })
			if err != nil {
				log.Fatalf("Could not start job of id: %v", v.ID)
			}

			updateJob := db.Job{
				ID:      v.ID,
				JobID:   int(id),
				LastRun: v.NextRun,
				NextRun: c.Entry(id).Next,
			}
			log.Printf("updating db for restarted job %v- job id %v, name: %s", v.ID, int(id), updateJob.JobName)
			go db.UpdateJobByID(&updateJob)
		}(job, v)

	}
	wg.Wait()
	//db update pending
	log.Println("All jobs restarted")
}
