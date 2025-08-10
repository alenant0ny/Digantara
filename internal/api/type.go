package api

import "encoding/json"

type CreateJobRequest struct {
	Cron string          `json:"cron"`
	Type string          `json:"type"` // SMS, Email, etc.
	Data json.RawMessage `json:"data"` // Raw so we can decode later based on Type
}
