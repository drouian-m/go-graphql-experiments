package post_api

import (
	post_service "github.com/drouian-m/go-graphql-experiments/internal/post"
	"github.com/graphql-go/graphql"
)

var postType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
		"create_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

type PostGraphqlController struct {
	postService *post_service.PostService
}

func NewPostGraphqlController(postService *post_service.PostService, query *graphql.Object, mutation *graphql.Object) {
	controller := PostGraphqlController{
		postService: postService,
	}

	query.AddFieldConfig("posts", controller.listPosts())
	query.AddFieldConfig("post", controller.getPost())
	mutation.AddFieldConfig("createPost", controller.createPost())
}

func (pgc *PostGraphqlController) listPosts() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(postType),
		Description: "List of posts",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			posts := pgc.postService.ListPosts()
			return posts, nil
		},
	}
}

func (pgc *PostGraphqlController) getPost() *graphql.Field {
	return &graphql.Field{
		Type:        postType,
		Description: "Get post",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(string)
			post, err := pgc.postService.GetPost(id)
			if err != nil {
				return nil, err
			}
			return post, nil
		},
	}
}

func (pgc *PostGraphqlController) createPost() *graphql.Field {
	return &graphql.Field{
		Type:        postType,
		Description: "Create new post",
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"content": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"author": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			title, _ := params.Args["title"].(string)
			content, _ := params.Args["content"].(string)
			author, _ := params.Args["author"].(string)

			post := pgc.postService.CreatePost(title, content, author)
			return post, nil
		},
	}

}
