package scheduler

import "fmt"

type Job interface {
	Run()
}

var jobRegistry = map[string]func() Job{
	"sms":   func() Job { return SMSJob{} },
	"email": func() Job { return EmailJob{} },
}

func GetJob(name string) (Job, error) {
	if constructor, ok := jobRegistry[name]; ok {
		return constructor(), nil
	}
	return nil, fmt.Errorf("unknown job type: %s", name)
}

type SMSJob struct{}

type EmailJob struct{}

func (s SMSJob) Run() {
	fmt.Println("Running SMS...")
}

func (e EmailJob) Run() {
	fmt.Println("Running Email...")
}
