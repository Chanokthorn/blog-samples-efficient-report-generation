package internal

import (
	"encoding/json"

	"github.com/Chanokthorn/blog-samples-efficient-report-generation/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type JobPublisher struct {
	channel *amqp.Channel
}

func NewJobPublisher(channel *amqp.Channel) *JobPublisher {
	return &JobPublisher{channel: channel}
}

func (jp *JobPublisher) PublishJob(jobMessage domain.JobMessage) error {
	msg, err := json.Marshal(jobMessage)
	if err != nil {
		return err
	}

	err = jp.channel.Publish("", "job", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	})
	return err
}
