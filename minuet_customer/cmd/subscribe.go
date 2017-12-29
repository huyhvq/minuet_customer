package cmd

import (
	"github.com/huyhvq/minuet_customer/subscribers"
	"github.com/lileio/pubsub"
	"github.com/spf13/cobra"
)

var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to and process queue messages",
	Run: func(cmd *cobra.Command, args []string) {
		pubsub.Subscribe(&subscribers.MinuetCustomerServiceSubscriber{})
	},
}

func init() {
	RootCmd.AddCommand(subscribeCmd)
}
