package cmd

import (
	"github.com/kuritka/plugin/runner"
	"github.com/kuritka/plugin/status"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "k8gb status",
	//TODO: long description
	//Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		status := status.New()
		runner.New(status).MustRun()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
