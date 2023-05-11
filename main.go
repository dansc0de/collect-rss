package main

import (
	"github.com/dansc0de/collect-rss/cmd/feed"
	"github.com/dansc0de/collect-rss/internal/util/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(feed.Cmd())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Log.Fatal(err)
	}
}
