package api

type CreateJobRequest struct {
	Cron    string `json:"cron"`
	Type    string `json:"type"` // SMS, Email, etc.
	Message string `json:"message"`
}
