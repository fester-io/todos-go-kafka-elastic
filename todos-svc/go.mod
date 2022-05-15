module github.com/fester-io/todos-todos-svc

go 1.18

replace github.com/fester-io/todos-lib => ../lib

require (
	github.com/fester-io/todos-lib v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.4.0
	github.com/twmb/franz-go v1.5.2
	go.uber.org/zap v1.21.0
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/klauspost/compress v1.15.4 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/twmb/franz-go/pkg/kmsg v1.0.0 // indirect
	github.com/twmb/go-rbtree v1.0.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
)
