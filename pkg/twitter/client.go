package twitter

import (
	"github.com/ChimeraCoder/anaconda"
)

type Client struct {
	*anaconda.TwitterApi
}

func New(consumerKey, consumerSecret, accessToken, accessTokenSecret string) *Client {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	return &Client{api}
}

func (c *Client) Tweet(message string) error {
	_, err := c.PostTweet(message, nil)
	return err
}
