package internal

import (
	"github.com/redis/go-redis/v9"
)

type JobDoneManager struct {
	pubSub          *redis.PubSub
	jobDoneRegistry map[string]chan struct{}
}

func NewJobDoneManager(pubSub *redis.PubSub) *JobDoneManager {
	return &JobDoneManager{
		pubSub:          pubSub,
		jobDoneRegistry: make(map[string]chan struct{}),
	}
}

func (jdm *JobDoneManager) Register(jobID string) (done <-chan struct{}, cancel func()) {
	doneChan := make(chan struct{})
	jdm.jobDoneRegistry[jobID] = doneChan
	cancel = func() {
		close(doneChan)
		delete(jdm.jobDoneRegistry, jobID)
	}
	return doneChan, cancel
}

func (jdm *JobDoneManager) Start() {
	go func() {
		for msg := range jdm.pubSub.Channel() {
			jobID := msg.Payload
			if done, ok := jdm.jobDoneRegistry[jobID]; ok {
				close(done)
				delete(jdm.jobDoneRegistry, jobID)
			}
		}
	}()
}
