package domain

type JobMessage struct {
	JobID        string `json:"job_id"`
	PreviousDays uint64 `json:"previous_days"`
}
