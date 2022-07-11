package tweet_service

import (
	"fmt"
	"github.com/drouian-m/array-utils"
	"github.com/google/uuid"
	"time"
)

type Tweet struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	Likes     int64     `json:"likes"`
}

type TweetService struct {
	tweets []Tweet
}

func NewTweetService() TweetService {
	return TweetService{
		tweets: []Tweet{},
	}
}

// ListTweets get all tweets
func (ts *TweetService) ListTweets() []Tweet {
	return ts.tweets
}

func (ts *TweetService) GetTweet(tweetId string) (*Tweet, error) {
	tweet := array.NewArray(ts.tweets).Find(func(t Tweet) bool {
		return t.ID == tweetId
	})

	if tweet.ID == "" {
		return nil, fmt.Errorf("GetTweetError - Tweet %s not found", tweetId)
	}

	return &tweet, nil
}

// CreateTweet create new tweet
func (ts *TweetService) CreateTweet(message string, author string) *Tweet {
	id, _ := uuid.NewUUID()
	tweet := Tweet{
		ID:        id.String(),
		Message:   message,
		Author:    author,
		CreatedAt: time.Now(),
		Likes:     0,
	}
	ts.tweets = append(ts.tweets, tweet)
	return &tweet
}

// LikeTweet like a tweet
func (ts *TweetService) LikeTweet(tweetId string) (*Tweet, error) {
	index := array.NewArray(ts.tweets).FindIndex(func(t Tweet) bool {
		return t.ID == tweetId
	})

	if index == -1 {
		return nil, fmt.Errorf("LikeTweetError - Tweet %s not found", tweetId)
	}
	ts.tweets[index].Likes += 1
	return &ts.tweets[index], nil
}
