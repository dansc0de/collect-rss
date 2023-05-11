package local

import (
	"github.com/spf13/cobra"
  "github.com/dansc0de/collect-rss/internal/pkg/util/logger"
)

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

