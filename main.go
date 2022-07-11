package main

import (
	"context"
	"errors"
	post_api "github.com/drouian-m/go-graphql-experiments/api/post"
	tweet_api "github.com/drouian-m/go-graphql-experiments/api/tweet"
	post_service "github.com/drouian-m/go-graphql-experiments/internal/post"
	tweet_service "github.com/drouian-m/go-graphql-experiments/internal/tweet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func graphqlHandler(schema *graphql.Schema) gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   schema,
		GraphiQL: true,
		Pretty:   true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphqlRouter(router *gin.Engine, query *graphql.Object, mutation *graphql.Object) {

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})

	router.GET("/graphql", graphqlHandler(&schema))
	router.POST("/graphql", graphqlHandler(&schema))
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	tweetService := tweet_service.NewTweetService()
	postService := post_service.NewPostService()
	tweet_api.NewTweetRestController(router, &tweetService)

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: graphql.Fields{},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: graphql.Fields{},
	})

	tweet_api.NewTweetGraphqlController(&tweetService, rootQuery, rootMutation)
	post_api.NewPostGraphqlController(&postService, rootQuery, rootMutation)

	graphqlRouter(router, rootQuery, rootMutation)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logrus.Error(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatal("Server forced to shutdown")
	}

	logrus.Info("Server exiting.")
}
