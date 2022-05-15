/*
Copyright (C) 2022 Jan te Beest <jtb@fester.io>
*/

package kafka

import (
	"context"
	"github.com/fester-io/todos-lib/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.uber.org/zap"
)

func StartConsumer() {
	log.Debug("starting consumer")
	go func() {
		for {
			fetches := cl.client.PollRecords(context.Background(), 10)
			if err := fetches.Err(); err != nil {
				log.Error("failed to fetch records", zap.Any("err", err))
				return
			}

			fetches.EachRecord(func(record *kgo.Record) {
				key := string(record.Key)
				log.Debug("record fetched", zap.String("key", key), zap.Any("record", record))

				produceRecord := &kgo.Record{
					Key:   record.Key,
					Topic: "todos-snd",
					Value: []byte("{}"),
				}

				cl.client.Produce(context.Background(), produceRecord, func(record *kgo.Record, err error) {
					if err != nil {
						log.Error("failed to produce record", zap.Any("record", produceRecord))
						return
					}

					log.Debug("record produced", zap.Any("record", produceRecord))

					//log.Debug("message produced, committing offsets", zap.Any("record", produceRecord))
					//if err = Commit(); err != nil {
					//	log.Error("failed to commit uncommitted offsets", zap.Any("err", err))
					//	return
					//}
				})
			})
		}
	}()
}
