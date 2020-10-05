package rss

import (
	"log"
	"sync"
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
	Published time.Time
}

func (*RSSReader) Read(urls []string) []*Feed {
	res := make([]*Feed, 0)
	wg := &sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			fp := gofeed.NewParser()
			feed, err := fp.ParseURL(u)
			if err != nil {
				log.Println(err)
				return
			}
			for _, item := range feed.Items {
				res = append(res, &Feed{Title: item.Title, Link: item.Link, Published: *item.PublishedParsed})
			}
		}(url)
	}
	wg.Wait()

	return res
}
