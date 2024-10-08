package main

import (
	"context"

	"github.com/Chanokthorn/blog-samples-efficient-report-generation/internal"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	server := gin.New()

	rabbitMQConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer rabbitMQConn.Close()

	rabbitMQChannel, err := rabbitMQConn.Channel()
	if err != nil {
		panic(err)
	}

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	jobPublisher := internal.NewJobPublisher(rabbitMQChannel)
	jobRepository := internal.NewJobRepository(mongoClient.Database("report").Collection("job"))

	apiHandler := internal.NewAPIHandler(jobPublisher, jobRepository)

	server.POST("/", apiHandler.GenerateReport)
	server.GET("/:job_id", apiHandler.GetReport)

	server.Run(":3000")
}
