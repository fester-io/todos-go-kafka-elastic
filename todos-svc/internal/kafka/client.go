package kafka

import (
	"context"
	"github.com/fester-io/todos-lib/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.uber.org/zap"
)

type client struct {
	client *kgo.Client
}

var cl client

// InitClient initializes the kafka client
func InitClient() error {
	cl1, err := kgo.NewClient(
		kgo.SeedBrokers("redpanda.redpanda-system:9092"),
		kgo.ConsumerGroup("todos-svc"),
		kgo.ConsumeTopics("todos-rcv"),
		//kgo.DisableAutoCommit(),
	)
	if err != nil {
		return err
	}

	cl = client{cl1}
	return nil
}

// Commit commits uncommitted offsets
func Commit() error {
	log.Debug("uncommitted offsets", zap.Any("offsets", cl.client.UncommittedOffsets()))
	if err := cl.client.CommitUncommittedOffsets(context.Background()); err != nil {
		return err
	}

	return nil
}
