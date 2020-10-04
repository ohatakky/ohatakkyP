package blog

import (
	"log"

	"github.com/ohatakky/ohatakkyp/pkg/rss"
)

var (
	urls = []string{
		"https://future-architect.github.io/atom.xml",
		"https://buildersbox.corp-sansan.com/rss",
	}
)

func init() {
	// todo: init config
}

func Exec() {
	reader := rss.New()
	feeds := reader.Read(urls)
	for _, feed := range feeds {
		log.Println(feed)
		// todo: 前回バッチ実時時間以降のフィードをツイート
	}

	// 	ctx := context.Background()
	// 	scheduler, err := schedule.New(ctx, os.Getenv("GCP_PROJECT"), os.Getenv("GCP_REGION_SCHEDULER"))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	job, err := scheduler.PrevJob(ctx)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(job)
}