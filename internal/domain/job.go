package domain

type Job struct {
	ID           string `json:"id"`
	PreviousDays uint64 `json:"previous_days"`
	Done         bool   `json:"done"`
	Content      Report `json:"content"`
}
