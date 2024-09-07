package internal

import (
	"strconv"

	"github.com/Chanokthorn/blog-samples-efficient-report-generation/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type APIHandler struct {
	// reportGenerator *ReportGenerator <<< removed from api handler, used by consumer instead
	jobPublisher  *JobPublisher
	jobRepository *JobRepository
}

func NewAPIHandler(
	jobPublisher *JobPublisher,
	jobRepository *JobRepository,
) *APIHandler {
	return &APIHandler{
		jobPublisher:  jobPublisher,
		jobRepository: jobRepository,
	}
}

func (h *APIHandler) GenerateReport(c *gin.Context) {
	previousDays, err := strconv.Atoi(c.Query("previous_days"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid previous_days parameter"})
		return
	}

	// report, err := ah.reportGenerator.GenerateReport(uint64(previousDays))
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "failed to generate report"})
	// 	return
	// }

	jobID := uuid.New().String()

	// insert a job to database for later update
	job := domain.Job{
		ID:           jobID,
		PreviousDays: uint64(previousDays),
		Done:         false,
	}

	err = h.jobRepository.InsertJob(job)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to insert job"})
	}

	// instead of generating the report, publish a job to the queue
	err = h.jobPublisher.PublishJob(jobID)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to publish job"})
		return
	}

	c.JSON(200, gin.H{"jobID": jobID})
}

func (h *APIHandler) GetReport(c *gin.Context) {
	// get job from database
	jobID := c.Param("job_id")
	job, err := h.jobRepository.FindJobByID(jobID)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to find job"})
		return
	}

	// return report if job is done
	if job.Done {
		c.JSON(200, gin.H{"report": job.Content})
		return
	}

	c.JSON(200, gin.H{"status": "processing"})
}
