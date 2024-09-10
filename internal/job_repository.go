package internal

import (
	"context"

	"github.com/Chanokthorn/blog-samples-efficient-report-generation/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobRepository struct {
	collection *mongo.Collection
}

func NewJobRepository(collection *mongo.Collection) *JobRepository {
	return &JobRepository{collection: collection}
}

func (jr *JobRepository) InsertJob(job domain.Job) error {
	_, err := jr.collection.InsertOne(context.Background(), job)
	return err
}

func (jr *JobRepository) FindJobByID(id string) (domain.Job, error) {
	var job domain.Job
	err := jr.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&job)
	return job, err
}

func (jr *JobRepository) UpdateJobDone(id string, report any) error {
	_, err := jr.collection.UpdateOne(context.Background(), bson.M{"id": id}, bson.M{"$set": bson.M{"done": true, "report": report}})
	return err
}
