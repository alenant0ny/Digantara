package scheduler

import "fmt"

type Job interface {
	Run(message string)
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

func (s SMSJob) Run(message string) {
	sms := "SMS - " + message
	fmt.Println(sms)
}

func (e EmailJob) Run(message string) {
	email := "Email - " + message
	fmt.Println(email)
}
