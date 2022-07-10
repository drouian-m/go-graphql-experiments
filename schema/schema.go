package schema

import (
	"time"

	"github.com/graphql-go/graphql"
)

var TweetList []Tweet

type Tweet struct {
	ID        string
	Message   string
	Author    string
	CreatedAt time.Time
	Likes     int64
}

// define custom GraphQL ObjectType `todoType` for our Golang struct `Todo`
// Note that
// - the fields in our todoType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
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

// root query
// we just define a trivial example here, since root query is required.
// Test with curl
// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{

		/*
		   curl -g 'http://localhost:8080/graphql?query={todo(id:"b"){id,text,done}}'
		*/
		"tweet": &graphql.Field{
			Type:        tweetType,
			Description: "Get single todo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				idQuery, isOK := params.Args["id"].(string)
				if isOK {
					// Search for el with id
					for _, tweet := range TweetList {
						if tweet.ID == idQuery {
							return tweet, nil
						}
					}
				}

				return Tweet{}, nil
			},
		},

		/*
		   curl -g 'http://localhost:8080/graphql?query={todoList{id,text,done}}'
		*/
		"tweets": &graphql.Field{
			Type:        graphql.NewList(tweetType),
			Description: "List of tweets",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return TweetList, nil
			},
		},
	},
})

// define schema, with our rootQuery and rootMutation
var TweetSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
