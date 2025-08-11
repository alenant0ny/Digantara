# Digantara Scheduler

## Setup
1. Clone the **Digantara** repository  
2. At root (`Digantara`) run:
   1. `chmod +x wait-for.sh`
   2. `docker-compose up --build`

---

## Endpoints

### 1. **GET** `/jobs`
List all jobs from the database.

---

### 2. **POST** `/jobs`

**JSON Input**:
```json
{
   "cron": "*/10 * * * * *",          // required
   "type": "sms",                     // required
   "message": "cron task message",
   "name": ""
}
```
As of now, the scheduler only accepts cron expressions from the input. This needs to be updated to accept higher-level information like "everyday", "monday", "2pm", etc., and translate it into a cron expression in the program logic.
This endpoint creates a record in the database for newly created jobs.
```json
{
   "cron": "*/10 * * * * *",         //every ten second
   "cron": "* */10 * * * *",         //every ten minutes
   "cron": "0 10 14 * * 1"            //every monday at 2.10pm
}
```
```
*   *   *   *   *   *
│   │   │   │   │   └─ day of week (0–6 or SUN–SAT)
│   │   │   │   └──── month (1–12 or JAN–DEC)
│   │   │   └──────── day of month (1–31)
│   │   └──────────── hour (0–23)
│   └──────────────── minute (0–59)
└──────────────────── second (0–59)
```

---

### 3. **GET** `/jobs/id/:id`
Get details of a job by id from database

---

## Project Startup
On startup, the app schedules all jobs available in the db, so that it acts like restarting jobs that were not run due to the application downtime. The db will be updated accordingly.

## Adding New Job
1. Create a struct for new job with Run method, so that it implements Job interface in internal/scheduler/jobs.go.
2. In internal/scheduler/jobs.go, modify variable jobRegistry with new job key.
```
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
```
```
var jobRegistry = map[string]func() Job{
	"sms":   func() Job { return SMSJob{} },
	"email": func() Job { return EmailJob{} },
}
```
