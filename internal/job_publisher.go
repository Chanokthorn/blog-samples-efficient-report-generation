package internal

import amqp "github.com/rabbitmq/amqp091-go"

type JobPublisher struct {
	channel *amqp.Channel
}

func NewJobPublisher(channel *amqp.Channel) *JobPublisher {
	return &JobPublisher{channel: channel}
}

func (jp *JobPublisher) PublishJob(jobID string) error {
	err := jp.channel.Publish("", "job", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(jobID),
	})
	return err
}
