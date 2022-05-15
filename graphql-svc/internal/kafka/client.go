package kafka

import (
	"github.com/fester-io/todos-lib/rand"
	"github.com/twmb/franz-go/pkg/kgo"
)

type client struct {
	client   *kgo.Client
	requests map[string]chan *kgo.Record
}

var cl client

func InitClient() error {
	cl1, err := kgo.NewClient(
		kgo.SeedBrokers("redpanda.redpanda-system:9092"),
		kgo.ConsumerGroup("graphql-svc-"+rand.ShortString(4)),
		kgo.ConsumeTopics("todos-snd", "users-snd"),
	)
	if err != nil {
		return err
	}

	cl = client{cl1, make(map[string]chan *kgo.Record)}
	return nil
}
