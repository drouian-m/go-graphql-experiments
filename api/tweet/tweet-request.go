package tweet_api

import (
	"time"
)

type TweetRequest struct {
	Message string `json:"message"`
	Author  string `json:"author"`
}

type TweetResponse struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	Author    string    `json:"author"`
	Likes     int64     `json:"likes"`
}
