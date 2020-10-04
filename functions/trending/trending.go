package trending

import (
	"os"
	"time"

	"github.com/ohatakky/ohatakkyp/pkg/trending"
	"github.com/ohatakky/ohatakkyp/pkg/tweet"
)

func Exec() error {
	cli := trending.New()
	items, err := cli.Read()
	if err != nil {
		return err
	}

	twitter := tweet.New(os.Getenv("API_KEY"), os.Getenv("API_SECRET"), os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	for i, item := range items {
		// note: note: avoid limited on the twitter API
		if i >= 5 {
			break
		}
		time.Sleep(5 * time.Second)

		err := twitter.Tweet(item.Link)
		if err != nil {
			continue
		}
	}

	return nil
}
