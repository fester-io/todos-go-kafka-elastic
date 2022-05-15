package kafka

import (
	"context"
	"github.com/fester-io/todos-lib/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.uber.org/zap"
)

// TODO we want to commit manually and only when the operation succeeds

func StartConsumer() {
	log.Debug("starting consumer")
	go func() {
		// TODO make this stop gracefully!
		for {
			fetches := cl.client.PollRecords(context.Background(), 10)
			if err := fetches.Err(); err != nil {
				log.Error("failed to fetch records", zap.Any("err", err))
				return
			}

			fetches.EachRecord(func(record *kgo.Record) {
				key := string(record.Key)
				log.Debug("record fetched", zap.String("key", key), zap.String("val", string(record.Value)))
				ch, ok := cl.requests[key]
				if ok {
					log.Debug("request found", zap.Any("key", key))
					delete(cl.requests, key)
					ch <- record
				} else {
					log.Debug("request not found", zap.String("key", key))
				}
			})
		}
	}()
}
