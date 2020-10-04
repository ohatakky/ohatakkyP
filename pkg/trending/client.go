package trending

import "net/http"

// curl https://ghapi.huchen.dev/repositories?since=daily | jq .
const endpoint = "https://ghapi.huchen.dev/repositories?since=daily"

type Client struct{}

func New() *Client {
	return &Client{}
}

type Item struct {
	Title string
	Link  string
}

func (*Client) Read() ([]*Item, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return nil, nil
}
