package rss

import (
	"context"
	"fmt"
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

func (*RSSReader) Read(ctx context.Context, urls []string) []*Feed {
	res := make([]*Feed, 0)
	wg := &sync.WaitGroup{}

	fmt.Println("-------- start rss reader ------------")
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			fp := gofeed.NewParser()
			feed, err := fp.ParseURLWithContext(u, ctx)
			if err != nil {
				log.Println("error: occured by parseURL(", u, ") >>> ", err)
				return
			}
			for _, item := range feed.Items {
				res = append(res, &Feed{Title: item.Title, Link: item.Link, Published: *item.PublishedParsed})
			}
		}(url)
	}
	wg.Wait()
	fmt.Println("-------- done rss reader ------------")

	return res
}
