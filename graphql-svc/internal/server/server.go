package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fester-io/todos-gql-kafka/graph"
	"github.com/fester-io/todos-gql-kafka/graph/generated"
	"github.com/fester-io/todos-lib/log"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func Start(port int64) error {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router := chi.NewRouter()
	router.Use(JwtAuth)
	router.Handle("/graphiql", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Infof("connect to http://localhost:%d/graphiql for GraphQL playground", port)
	if err := http.ListenAndServe("0.0.0.0:"+strconv.FormatInt(port, 10), router); err != nil {
		log.Fatal("failed to start server", zap.Any("err", err))
	}

	log.Info("server stopped")
	return nil
}
