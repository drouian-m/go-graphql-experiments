package post_service

import (
	"fmt"
	"github.com/drouian-m/array-utils"
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
}

type PostService struct {
	posts []Post
}

func NewPostService() PostService {
	return PostService{
		posts: []Post{},
	}
}

// ListPosts get all posts
func (ps *PostService) ListPosts() []Post {
	return ps.posts
}

func (ps *PostService) GetPost(postId string) (*Post, error) {
	post := array.NewArray(ps.posts).Find(func(p Post) bool {
		return p.ID == postId
	})

	if post.ID == "" {
		return nil, fmt.Errorf("GetPostError - Post %s not found", postId)
	}

	return &post, nil
}

// CreatePost create new post
func (ps *PostService) CreatePost(title string, content string, author string) *Post {
	id, _ := uuid.NewUUID()
	post := Post{
		ID:        id.String(),
		Title:     title,
		Content:   content,
		Author:    author,
		CreatedAt: time.Now(),
	}
	ps.posts = append(ps.posts, post)
	return &post
}
