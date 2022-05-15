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
	// FIXME by generating a new groupid every time this service starts,
	// it effectively resets the offset and will read all records from 0.
	// Although they will not be handled (we need to fix this too though,
	// user_id is not what we should store in the map, use something else),
	// we should determine the previous offset or use a somewhat static node_id.
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
