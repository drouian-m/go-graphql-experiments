package tweet_api

import (
	tweet_service "github.com/drouian-m/go-graphql-experiments/internal/tweet"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TweetRestController struct {
	tweetService *tweet_service.TweetService
}

func NewTweetRestController(router *gin.Engine, tweetService *tweet_service.TweetService) {
	controller := TweetRestController{
		tweetService: tweetService,
	}

	router.GET("/tweets", controller.listTweets)
	router.POST("/tweets", controller.createTweet)
	router.POST("/tweets/:id/like-tweet", controller.likeTweet)
}

func (tgc *TweetRestController) listTweets(c *gin.Context) {
	tweets := tgc.tweetService.ListTweets()
	c.JSON(http.StatusOK, tweets)
}

func (tgc *TweetRestController) createTweet(c *gin.Context) {
	var input TweetRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tweet := tgc.tweetService.CreateTweet(input.Message, input.Author)
	c.JSON(http.StatusCreated, tweet)

}

func (tgc *TweetRestController) likeTweet(c *gin.Context) {
	tweetID := c.Param("id")
	tweet, err := tgc.tweetService.LikeTweet(tweetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tweet)
}
