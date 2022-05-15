module github.com/fester-io/todos-gql-kafka

go 1.18

replace github.com/fester-io/todos-lib => ../lib

require (
	github.com/99designs/gqlgen v0.17.5
	github.com/fester-io/todos-lib v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.7
	github.com/spf13/cobra v1.4.0
	github.com/twmb/franz-go v1.5.2
	github.com/vektah/gqlparser/v2 v2.4.2
	go.uber.org/zap v1.21.0
)

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/klauspost/compress v1.15.4 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/twmb/franz-go/pkg/kmsg v1.0.0 // indirect
	github.com/twmb/go-rbtree v1.0.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
)
