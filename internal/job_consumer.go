package internal

import (
	"encoding/json"
	"log"

	"github.com/Chanokthorn/blog-samples-efficient-report-generation/internal/domain"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	channel          *amqp.Channel
	jobRepository    *JobRepository
	reportGenerator  *ReportGenerator
	jobDonePublisher *JobDonePublisher
}

func NewConsumer(
	channel *amqp.Channel,
	jobRepository *JobRepository,
	reportGenerator *ReportGenerator,
	jobDonePublisher *JobDonePublisher,
) *Consumer {
	return &Consumer{
		channel:          channel,
		jobRepository:    jobRepository,
		reportGenerator:  reportGenerator,
		jobDonePublisher: jobDonePublisher,
	}
}

func (c *Consumer) Consume() {
	log.Println("Start consuming jobs...")

	msgs, err := c.channel.Consume("job", "", true, true, false, false, nil)
	if err != nil {
		log.Fatalf("failed to consume: %v", err)
		return
	}

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			// parse job message
			var jobMessage domain.JobMessage
			err := json.Unmarshal(msg.Body, &jobMessage)
			if err != nil {
				log.Printf("failed to parse job message: %v", err)
				continue
			}

			// generate report based on job
			report, err := c.reportGenerator.GenerateReport(jobMessage.PreviousDays)
			if err != nil {
				log.Printf("failed to generate report: %v", err)
				continue
			}

			// update job status to done with the report
			err = c.jobRepository.UpdateJobDone(jobMessage.JobID, report)
			if err != nil {
				log.Printf("failed to update job: %v", err)
				continue
			}

			// publish job done event
			err = c.jobDonePublisher.PublishJobDone(jobMessage.JobID)
			if err != nil {
				log.Printf("failed to publish job done: %v", err)
				continue
			}
		}
	}()

	<-forever
}
