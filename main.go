package main

import (
	"github.com/spf13/cobra"
	"github.com/dansc0de/collect-rss/cmd/feed"
	"github.com/dansc0de/collect-rss/internal/utils/logger"
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
