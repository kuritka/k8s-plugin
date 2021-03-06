package cmd

import (
	"github.com/kuritka/plugin/cmd/internal/runner"
	"github.com/kuritka/plugin/common/guard"
	k8sctx2 "github.com/kuritka/plugin/common/k8sctx"
	"github.com/kuritka/plugin/status"
	"github.com/spf13/cobra"
)

var statusOptions status.Options

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "k8gb status",
	//TODO: long description
	//Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		statusOptions.Context, err = k8sctx2.NewContextFactory(args).Get()
		guard.FailOnError(err, "error when building command context")
		status := status.New(statusOptions)
		runner.New(status).MustRun()
	},
}

func init() {
	//TODO: fix description
	statusCmd.Flags().StringVarP(&statusOptions.Namespace, "namespace", "n", "default", "k8gb namespace")
	err := statusCmd.MarkFlagRequired("namespace")
	guard.FailOnError(err, "namespace required")
	rootCmd.AddCommand(statusCmd)
}
