package tweet_api

import (
	tweet_service "github.com/drouian-m/go-graphql-experiments/internal/tweet"
	"github.com/graphql-go/graphql"
)

var tweetType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Tweet",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
		"create_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"likes": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

type TweetGraphqlController struct {
	tweetService *tweet_service.TweetService
	TweetSchema  graphql.Schema
}

func NewTweetGraphqlController(tweetService *tweet_service.TweetService) TweetGraphqlController {

	var TweetQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "TweetQuery",
		Fields: graphql.Fields{
			"tweets": &graphql.Field{
				Type:        graphql.NewList(tweetType),
				Description: "List of tweets",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					tweets := tweetService.ListTweets()
					return tweets, nil
				},
			},
		},
	})

	var TweetMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "TweetMutation",
		Fields: graphql.Fields{
			"createTweet": &graphql.Field{
				Type:        tweetType, // the return type for this field
				Description: "Create new tweet",
				Args: graphql.FieldConfigArgument{
					"message": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"author": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					message, _ := params.Args["message"].(string)
					author, _ := params.Args["author"].(string)

					tweet := tweetService.CreateTweet(message, author)
					return tweet, nil
				},
			},
			"likeTweet": &graphql.Field{
				Type:        tweetType, // the return type for this field
				Description: "Like a tweet",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)
					tweet, err := tweetService.LikeTweet(id)
					if err != nil {
						return nil, err
					}
					return tweet, nil
				},
			},
		},
	})

	var tweetSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    TweetQuery,
		Mutation: TweetMutation,
	})

	controller := TweetGraphqlController{
		tweetService: tweetService,
		TweetSchema:  tweetSchema,
	}

	return controller
}
