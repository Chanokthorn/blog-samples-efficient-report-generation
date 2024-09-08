package internal

import (
	"log"

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
			// get job from database
			job, err := c.jobRepository.FindJobByID(string(msg.Body))
			if err != nil {
				log.Printf("failed to find job: %v", err)
				continue
			}

			// generate report based on job
			report, err := c.reportGenerator.GenerateReport(job.PreviousDays)
			if err != nil {
				log.Printf("failed to generate report: %v", err)
				continue
			}

			// update job status to done with the report
			err = c.jobRepository.UpdateJobDone(job.ID, report)
			if err != nil {
				log.Printf("failed to update job: %v", err)
				continue
			}

			// publish job done event
			err = c.jobDonePublisher.PublishJobDone(job.ID)
			if err != nil {
				log.Printf("failed to publish job done: %v", err)
				continue
			}
		}
	}()

	<-forever
}
