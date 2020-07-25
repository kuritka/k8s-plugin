// Package cmd implements three different commands
// - status
// - test
// - install
package cmd

import (
	"context"
	"k8s.io/client-go/tools/clientcmd/api"
	"os"

	"github.com/kuritka/plugin/common/log"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type commandLineContext struct{
	configFlags *genericclioptions.ConfigFlags

	resultingContext     *api.Context
	resultingContextName string

	rawConfig      api.Config
	listNamespaces bool
	args           []string

	genericclioptions.IOStreams
}

var (
	logger = log.Log
	rootContext, rootContextCancel = context.WithCancel(context.Background())
	commandContext = commandLineContext{
		configFlags: genericclioptions.NewConfigFlags(true),
		IOStreams: genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr},
	}
	//Verbose output
	 Verbose bool
)

var rootCmd = &cobra.Command{
	Short: "k8gb plugins",
	//TODO: Long description
	//Long:  `load balancer demo`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			logger.Error().Msg("No parameters included")
			_ = cmd.Help()
			os.Exit(0)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		logger.Info().Msg("done..")
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

//Execute runs concrete command
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
