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
}

func NewTweetGraphqlController(tweetService *tweet_service.TweetService, query *graphql.Object, mutation *graphql.Object) {
	controller := TweetGraphqlController{
		tweetService: tweetService,
	}

	query.AddFieldConfig("tweets", controller.listTweets())
	query.AddFieldConfig("tweet", controller.getTweet())
	mutation.AddFieldConfig("createTweet", controller.createTweet())
	mutation.AddFieldConfig("likeTweet", controller.likeTweet())
}

func (tgc *TweetGraphqlController) listTweets() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(tweetType),
		Description: "List of tweets",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			tweets := tgc.tweetService.ListTweets()
			return tweets, nil
		},
	}
}

func (tgc *TweetGraphqlController) getTweet() *graphql.Field {
	return &graphql.Field{
		Type:        tweetType,
		Description: "Get tweet",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(string)
			tweet, err := tgc.tweetService.GetTweet(id)
			if err != nil {
				return nil, err
			}
			return tweet, nil
		},
	}
}

func (tgc *TweetGraphqlController) createTweet() *graphql.Field {
	return &graphql.Field{
		Type:        tweetType,
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

			tweet := tgc.tweetService.CreateTweet(message, author)
			return tweet, nil
		},
	}

}

func (tgc *TweetGraphqlController) likeTweet() *graphql.Field {
	return &graphql.Field{
		Type:        tweetType,
		Description: "Like a tweet",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(string)
			tweet, err := tgc.tweetService.LikeTweet(id)
			if err != nil {
				return nil, err
			}
			return tweet, nil
		},
	}
}
