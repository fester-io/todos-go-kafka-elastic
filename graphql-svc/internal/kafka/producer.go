package kafka

import (
	"context"
	"github.com/fester-io/todos-lib/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.uber.org/zap"
)

func ProduceAndWait(record *kgo.Record) chan *kgo.Record {
	ch := make(chan *kgo.Record, 1)

	cl.client.Produce(context.Background(), record, func(r *kgo.Record, err error) {
		if err != nil {
			log.Error("produce failed", zap.Any("err", err))
			return
		}

		log.Debug("record produced", zap.Any("record", record))
		cl.requests[string(record.Key)] = ch
	})

	return ch
}
