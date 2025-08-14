package scheduler

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

type Job interface {
	Run(message string)
	SetID(id cron.EntryID)
}

var jobRegistry = map[string]func() Job{
	"sms":   func() Job { return &SMSJob{} },
	"email": func() Job { return &EmailJob{} },
}

func GetJob(name string) (Job, error) {
	if constructor, ok := jobRegistry[name]; ok {
		return constructor(), nil
	}
	return nil, fmt.Errorf("unknown job type: %s", name)
}

type SMSJob struct {
	ID cron.EntryID
}

type EmailJob struct {
	ID cron.EntryID
}

func (s SMSJob) Run(message string) {
	sms := "SMS - " + message
	fmt.Println(s.ID, ":", sms)
}

func (s *SMSJob) SetID(id cron.EntryID) {
	s.ID = id
}

func (e EmailJob) Run(message string) {
	email := "Email - " + message
	fmt.Println(e.ID, ":", email)
}

func (e *EmailJob) SetID(id cron.EntryID) {
	e.ID = id
}
