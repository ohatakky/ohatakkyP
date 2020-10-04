package blog

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/ohatakky/ohatakkyp/pkg/rss"
	"github.com/ohatakky/ohatakkyp/pkg/schedule"
	"github.com/ohatakky/ohatakkyp/pkg/tweet"
)

var (
	urls = []string{
		"https://www.cdatablog.jp/rss",
		"https://future-architect.github.io/atom.xml",
		"https://buildersbox.corp-sansan.com/rss",
		"https://developers-jp.googleblog.com/atom.xml",
		"https://techblog.zozo.com/rss",
		"https://www.m3tech.blog/feed",
		"https://engineering.mercari.com/blog/feed.xml",
		"https://developer.hatenastaff.com/rss",
	}
)

func init() {
	godotenv.Load()
}

func Exec() error {
	ctx := context.Background()
	scheduler, err := schedule.New(ctx, os.Getenv("GCP_PROJECT"), os.Getenv("GCP_REGION_SCHEDULER"))
	if err != nil {
		return err
	}
	job, err := scheduler.GetJob(ctx, os.Getenv("GCP_SCHEDULER_BLOG_ID"))
	if err != nil {
		return err
	}

	twitter := tweet.New(os.Getenv("API_KEY"), os.Getenv("API_SECRET"), os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))

	reader := rss.New()
	feeds := reader.Read(urls)
	feeds = feeds[:5] // note: avoid limited on the twitter API
	for _, feed := range feeds {
		if feed.Published.After(job.LastAttemptTime) {
			content := fmt.Sprintf("%s %s", feed.Title, feed.Link)
			err := twitter.Tweet(content)
			if err != nil {
				continue
			}
		}
	}

	return nil
}
