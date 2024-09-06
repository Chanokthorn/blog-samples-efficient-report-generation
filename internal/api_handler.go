package internal

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	ReportGenerator *ReportGenerator
}

func NewAPIHandler(rg *ReportGenerator) *APIHandler {
	return &APIHandler{ReportGenerator: rg}
}

func (ah *APIHandler) GetReport(c *gin.Context) {
	previousDays, err := strconv.Atoi(c.Query("previous_days"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid previous_days parameter"})
		return
	}
	
	report, err := ah.ReportGenerator.GenerateReport(uint64(previousDays))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate report"})
		return
	}

	c.JSON(200, report)
}