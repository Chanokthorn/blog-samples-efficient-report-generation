package internal

import (
	"context"
	"strconv"
	"time"

	"github.com/Chanokthorn/blog-samples-efficient-report-generation/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type APIHandler struct {
	// reportGenerator *ReportGenerator <<< removed from api handler, used by consumer instead
	jobPublisher   *JobPublisher
	jobRepository  *JobRepository
	jobDoneManager *JobDoneManager
}

func NewAPIHandler(
	jobPublisher *JobPublisher,
	jobRepository *JobRepository,
	jobDoneManager *JobDoneManager,
) *APIHandler {
	return &APIHandler{
		jobPublisher:   jobPublisher,
		jobRepository:  jobRepository,
		jobDoneManager: jobDoneManager,
	}
}

func (h *APIHandler) GenerateReport(c *gin.Context) {
	previousDays, err := strconv.Atoi(c.Query("previous_days"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid previous_days parameter"})
		return
	}

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

	done, cancelRegistration := h.jobDoneManager.Register(jobID)

	// create context with timeout to limit wait time for asynchronous response
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	select {
	// if job is done within 3 seconds, query result and send within api
	case <-done:
		// get job from database
		newJob, err := h.jobRepository.FindJobByID(jobID)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to find job"})
			return
		}
		c.JSON(200, gin.H{"report": newJob.Content})
	// if job is not done within 3 seconds, sends job ID for front end for later query
	case <-ctx.Done():
		cancelRegistration()
		c.JSON(200, gin.H{
			"status": "processing",
			"jobID":  jobID,
		})
	}
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
