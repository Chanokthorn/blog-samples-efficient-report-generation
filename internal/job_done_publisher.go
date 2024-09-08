package internal

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type JobDonePublisher struct {
	client *redis.Client
}

func NewJobDonePublisher(client *redis.Client) *JobDonePublisher {
	return &JobDonePublisher{client: client}
}

func (jdp JobDonePublisher) PublishJobDone(jobID string) error {
	err := jdp.client.Publish(context.Background(), "job_done", jobID).Err()
	if err != nil {
		return err
	}

	return nil
}
