package main

import (
	"fmt"
	"github.com/drouian-m/go-graphql-experiments/schema"
	"github.com/graphql-go/handler"
	"net/http"
	"time"
)

func init() {
	tweet1 := schema.Tweet{ID: "qwe", Message: "Hello world !", Author: "zig", CreatedAt: time.Now(), Likes: 0}
	tweet2 := schema.Tweet{ID: "sjd", Message: "Hi !", Author: "zig", CreatedAt: time.Now(), Likes: 3}
	tweet3 := schema.Tweet{ID: "dof", Message: "Pouet !", Author: "zig", CreatedAt: time.Now(), Likes: 213}
	schema.TweetList = append(schema.TweetList, tweet1, tweet2, tweet3)

}

func main() {
	h := handler.New(&handler.Config{
		Schema:   &schema.TweetSchema,
		GraphiQL: true,
		Pretty:   true,
	})

	http.Handle("/graphql", h)
	fmt.Println("🚀 server started and available on http://localhost:8080")
	fmt.Println("... graphql endpoint is available at http://localhost:8080/graphql")
	http.ListenAndServe(":8080", nil)
}
