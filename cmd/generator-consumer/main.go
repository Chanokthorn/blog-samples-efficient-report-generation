package main

import (
	"context"

	"github.com/Chanokthorn/blog-samples-efficient-report-generation/internal"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	rabbitMQConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer rabbitMQConn.Close()

	rabbitMQChannel, err := rabbitMQConn.Channel()
	if err != nil {
		panic(err)
	}

	_, err = rabbitMQChannel.QueueDeclare(
		"job", // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		panic(err)
	}

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	jobRepository := internal.NewJobRepository(mongoClient.Database("report").Collection("job"))
	jobDonePublisher := internal.NewJobDonePublisher(redisClient)
	reportGenerator := internal.NewReportGenerator()

	consumer := internal.NewConsumer(rabbitMQChannel, jobRepository, reportGenerator, jobDonePublisher)

	consumer.Consume()
}
