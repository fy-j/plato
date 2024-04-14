package cmd

import (
	"github.com/hardcore-os/plato/client"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand()
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client open",
	Run:   ClientHandle,
}

func ClientHandle(cmd *cobra.Command, args []string) {
	client.RunMain()
}
