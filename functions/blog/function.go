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
		// newsletter
		"https://coin98analytics.substack.com/feed",
		"https://todayindefi.substack.com/feed",
		"https://banklessdao.substack.com/feed",
		"https://newsletter.banklesshq.com/feed",
		"https://newsletter.thedefiant.io/feed",
		"https://ethhub.substack.com/feed",
		"https://ournetwork.substack.com/feed",
		"https://defiprime.substack.com/feed",
		"https://defiweekly.substack.com/feed",
		"https://aavenews.substack.com/feed",
		"https://yieldfarmer.substack.com/feed",
		"https://digitalnative.substack.com/feed",
		"https://rewilder.substack.com/feed",
		"https://axie.substack.com/feed",
		"https://cryptex.substack.com/feed",
		"https://weekinethereumnews.com/feed",
		"https://dunedigest.substack.com/feed",
		"https://epns.substack.com/feed",

		// personal newsletter
		"https://www.platformer.news/feed",
		"https://li.substack.com/feed",
		"https://midroni.substack.com/feed",
		"https://pizzaparty.substack.com/feed",
		"https://dovemountain.substack.com/feed",
		"https://davidphelps.substack.com/feed",
		"https://learnhax.substack.com/feed",
		"https://www.uncomfortableprofit.com/feed",
		"https://lootproject.substack.com/feed",
		"https://jarvislabs.substack.com/feed",
		"https://robdog.substack.com/feed",
		"https://creatoreconomy.so/feed",
		"https://www.notboring.co/feed",
		"https://kinjalshah.substack.com/feed",
		"https://nickwidmer.substack.com/feed",
		"https://simondlr.substack.com/feed",
		"https://thedailygwei.substack.com/feed",
		"https://newsletter.withtally.com/feed",
		"https://doseofdefi.substack.com/feed",
		"https://willywoo.substack.com/feed",
		"https://bspeak.substack.com/feed",
		"https://nobumei.substack.com/feed",
		"https://oji3.substack.com/feed",

		// blog
		"https://messari.io/rss",
		"https://insights.glassnode.com/rss",
		"https://blog.mycrypto.com/rss",
		"https://dappradar.com/blog/feed",
		"https://future.a16z.com/feed",
		"https://finematics.com/feed",
		"https://nftnow.com/feed",
		"https://defipulse.com/blog/feed",
		"https://medium.com/feed/pinata",
		"https://defirate.com/feed",
		"https://www.blocmates.com/blogmates?format=rss",
		"https://defipocket.jp/index.php/feed",

		// personal blog
		"https://vitalik.ca/feed.xml",
		"https://balajis.com/author/balajis/rss",

		// daily news
		"https://decrypt.co/feed",
	}
)

func Exec() error {
	ctx := context.Background()
	scheduler, err := schedule.New(ctx, os.Getenv("GCP_PROJECT"), os.Getenv("GCP_REGION_SCHEDULER"))
	if err != nil {
		return err
	}
	twitter := tweet.New(os.Getenv("API_KEY"), os.Getenv("API_SECRET"), os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	reader := rss.New()

	job, err := scheduler.GetJob(ctx, os.Getenv("GCP_SCHEDULER_BLOG_ID"))
	if err != nil {
		return err
	}
	feeds := reader.Read(ctx, urls)
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
