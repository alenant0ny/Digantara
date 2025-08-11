package api

type CreateJobRequest struct {
	Cron    string `json:"cron" validate:"required"`
	Type    string `json:"type" validate:"required"` // sms, email, etc.
	Message string `json:"message"`
	Name    string `json:"name"`
}
