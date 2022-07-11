package main

import (
	"context"
	"errors"
	tweet_api "github.com/drouian-m/go-graphql-experiments/api/tweet"
	tweet_service "github.com/drouian-m/go-graphql-experiments/internal/tweet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func graphqlHandler(tweetGraphqlController tweet_api.TweetGraphqlController) gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &tweetGraphqlController.TweetSchema,
		GraphiQL: true,
		Pretty:   true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	tweetService := tweet_service.NewTweetService()
	tweet_api.NewTweetRestController(router, &tweetService)
	tweetGraphqlController := tweet_api.NewTweetGraphqlController(&tweetService)
	router.GET("/graphql", graphqlHandler(tweetGraphqlController))
	router.POST("/graphql", graphqlHandler(tweetGraphqlController))

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
