package rss

import (
	"time"

	"github.com/mmcdole/gofeed"
)

type RSSReader struct{}

func New() *RSSReader {
	return &RSSReader{}
}

type Feed struct {
	Title     string
	Link      string
	Published *time.Time
}

func (*RSSReader) Read(urls []string) []*Feed {
	res := make([]*Feed, 0)
	for _, url := range urls {
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(url)
		if err != nil {
			continue
		}

		for _, item := range feed.Items {
			res = append(res, &Feed{Title: item.Title, Link: item.Link, Published: item.PublishedParsed})
		}
	}
	return res
}
