package internal

import (
	"time"

	"github.com/Chanokthorn/blog-samples-efficient-report-generation/internal/domain"
)


type ReportGenerator struct{}

func NewReportGenerator() *ReportGenerator {
	return &ReportGenerator{}
}

func (rg *ReportGenerator) GenerateReport(previousDays uint64) (domain.Report, error) {
	// simulate a delay based on the number of previous days
	time.Sleep(time.Duration(previousDays) * time.Second)

	switch {
	case previousDays < 3:
		return domain.Report{
			Title:   "Small Report",
			Content: "This is a small mock report.",
		}, nil
	default:
		return domain.Report{
			Title:   "Large Report",
			Content: "This is a large mock report with a lot of content...",
		}, nil
	}
}
