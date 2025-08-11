package api

type CreateJobRequest struct {
	Cron    string `json:"cron"`
	Type    string `json:"type"` // sms, email, etc.
	Message string `json:"message"`
	Name    string `json:"name"`
}
