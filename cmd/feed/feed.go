package feed

import (
	"github.com/dansc0de/collect-rss/internal/model"
	"github.com/dansc0de/collect-rss/internal/util/logger"
	"github.com/spf13/cobra"
)

var feedURLs = []string{
	"https://blog.golang.org/feed.atom",
	"http://feeds.kottke.org/main",
}

func fetchFeeds(urls []string) []*Feed {
	// Create a channel to process the feeds
	feedc := make(chan feedResult, len(feedURLs))

	// Start a goroutine for each feed url
	for _, u := range urls {
		go fetchFeed(u, feedc)
	}

	// Wait for the goroutines to write their results to the channel
	feeds := []*Feed{}
	for i := 0; i < len(urls); i++ {
		res := <-feedc
		// If the goroutine errors out, we'll just wait for others
		if res.err != nil {
			continue
		}
		feeds = append(feeds, res.feed)
	}

	return feeds
}

func fetchFeed(url string, feedc chan feedResult) {
	// Create a client with a default timeout
	net := &http.Client{
		Timeout: time.Second * 10,
	}
	// Issue a GET request for the feed
	res, err := net.Get(url)
	// If there was an error write that to the channel and return immediately
	if err != nil {
		feedc <- feedResult{nil, err}
		return
	}
	defer res.Body.Close()
	// Read the body of the request and parse the feed
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		feedc <- feedResult{nil, err}
		return
	}
	feed, err := parseFeed(body)
	if err != nil {
		feedc <- feedResult{nil, err}
		return
	}
	feedc <- feedResult{feed, nil}
}

func parseFeed(body []byte) (*Feed, error) {
	feed := Feed{}
	err := xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}
	return &feed, nil
}

var feedCmd = &cobra.Command{
	Use:   "feed",
	Short: "Collects rss feeds",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Log.Info("Generating output")
		return nil
	},
}

func Cmd() *cobra.Command {
	return feedCmd
}
