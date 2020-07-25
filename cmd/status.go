package cmd

import (
	"fmt"
	"github.com/kuritka/plugin/runner"
	"github.com/kuritka/plugin/status"
	"github.com/spf13/cobra"
)


var statusCmdOptions status.Options

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "k8gb status",
	//TODO: long description
	//Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		k8sContext,_ := commandContext.configFlags.ToRawKubeConfigLoader().RawConfig()
		fmt.Println(k8sContext)
		status := status.New(statusCmdOptions)
		runner.New(status).MustRun()
	},
}

func init() {
	//TODO: fix description
	statusCmd.Flags().StringVarP(&statusCmdOptions.Namespace, "namespace", "n", "default", "k8gb namespace")
	statusCmd.MarkFlagRequired("namespace")
	rootCmd.AddCommand(statusCmd)
}
