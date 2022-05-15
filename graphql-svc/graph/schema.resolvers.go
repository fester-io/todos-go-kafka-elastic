package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/fester-io/todos-gql-kafka/internal/kafka"
	"github.com/fester-io/todos-lib/log"
	"github.com/twmb/franz-go/pkg/kgo"
	"reflect"
	"unsafe"

	"github.com/fester-io/todos-gql-kafka/graph/generated"
	"github.com/fester-io/todos-gql-kafka/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Todos retrieves all the TODOs for the current user
// TODO get current user from context (somehow, use a middleware to extract the user ID from JWT?)
// TODO we will need the user and the roles, as the admin should be able to retrieve all, or use 2 different queries for that
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	record := &kgo.Record{
		Topic: "todos-rcv",
		Key:   []byte(ctx.Value("user_id").(string)),
		Value: []byte(`{}`)}

	ch := kafka.ProduceAndWait(record)

	ret := <-ch
	fmt.Printf("ret=%v\n", ret)

	var todos []*model.Todo

	return todos, nil
}

func printContext(ctx any, inner bool) {
	ctxValues := reflect.ValueOf(ctx).Elem()
	ctxKeys := reflect.TypeOf(ctx).Elem()

	if !inner {
		log.Infof("Fields for %s.%s", ctxKeys.PkgPath(), ctxKeys.Name())
	}

	if ctxKeys.Kind() == reflect.Struct {
		for i := 0; i < ctxKeys.NumField(); i++ {
			reflectVal := ctxValues.Field(i)
			reflectVal = reflect.NewAt(reflectVal.Type(), unsafe.Pointer(reflectVal.UnsafeAddr())).Elem()

			reflectField := ctxKeys.Field(i)
			if reflectField.Name == "Context" {
				printContext(reflectVal.Interface(), true)
			} else {
				fmt.Printf("\n%v: %v\n", reflectField.Name, reflectVal.Interface())
			}
		}
	} else {
		log.Infof("Context is empty")
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
