package scheduler

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

var c = cron.New(cron.WithSeconds())
var jobIDs = map[string]cron.EntryID{}

func StartScheduler() {
	c.Start()
}

func AddJob(name, cronExpr string) (cron.EntryID, error) {
	job, err := GetJob(name)
	if err != nil {
		return 0, err
	}
	id, err := c.AddFunc(cronExpr, func() { job.Run() })
	if err != nil {
		return 0, err
	}
	jobIDs[name] = id
	return jobIDs[name], nil
}

// func RemoveJob(name string) {
// 	if id, exists := jobIDs[name]; exists {
// 		c.Remove(id)
// 		delete(jobIDs, name)
// 	}
// }

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
