package blog

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ohatakky/ohatakkyp/pkg/rss"
	"github.com/ohatakky/ohatakkyp/pkg/schedule"
	"github.com/ohatakky/ohatakkyp/pkg/tweet"
)

var (
	urls = []string{
		// tech
		"https://future-architect.github.io/atom.xml",
		"https://buildersbox.corp-sansan.com/rss",
		"https://developers-jp.googleblog.com/atom.xml",
		"https://techblog.zozo.com/rss",
		"https://www.m3tech.blog/feed",
		"https://engineering.mercari.com/blog/feed.xml",
		"https://developer.hatenastaff.com/rss",
		"https://medium.com/feed/studist-dev",
		"https://eng.uber.com/feed",
		"https://medium.com/feed/airbnb-engineering",
		"https://medium.com/feed/google-cloud",

		// media
		"https://news.ycombinator.com/rss",
		"https://techcrunch.com/feed",
		"https://a16z.com/feed",
		"https://news.crunchbase.com/feed",

		// media
		"https://jp.techcrunch.com/feed",
		"https://techable.jp/feed",
		"https://www.businessinsider.jp/feed/index.xml",
		"https://techblitz.com/feed",

		// personal media
		"https://koheeiokubo.substack.com/feed",
		"https://note.com/okb777/rss",
		"https://note.com/0915hikaru/rss",
	}
)

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

	// todo: schedule queue
	for _, feed := range feeds {
		if feed.Published.Before(job.LastAttemptTime) {
			continue
		}

		err := twitter.Tweet(fmt.Sprintf("%s %s", feed.Title, feed.Link))
		if err != nil {
			continue
		}

		time.Sleep(10 * time.Second)
	}

	return nil
}
