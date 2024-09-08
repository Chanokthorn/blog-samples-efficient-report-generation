package internal

import (
	"time"
)

type Report struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ReportGenerator struct{}

func NewReportGenerator() *ReportGenerator {
	return &ReportGenerator{}
}

func (rg *ReportGenerator) GenerateReport(previousDays uint64) (Report, error) {
	// simulate a delay based on the number of previous days
	time.Sleep(time.Duration(previousDays) * time.Second)

	switch {
	case previousDays < 3:
		return Report{
			Title:   "Small Report",
			Content: "This is a small mock report.",
		}, nil
	default:
		return Report{
			Title:   "Large Report",
			Content: "This is a large mock report with a lot of content...",
		}, nil
	}
}
