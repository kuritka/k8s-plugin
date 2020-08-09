package k8sctx

import (
	"context"
	"os"

	"k8s.io/cli-runtime/pkg/genericclioptions"
)

//ContextFactory keeps k8s context
type ContextFactory struct {
	args []string
}

//NewContextFactory returns context of command
func NewContextFactory(args []string) *ContextFactory {
	return &ContextFactory{
		args: args,
	}
}

//Get returns context
func (cf *ContextFactory) Get() (*Context, error) {
	var err error
	ctx := new(Context)
	ctx.Command = new(Command)
	ctx.K8s = new(K8s)
	configFlags := genericclioptions.NewConfigFlags(true)
	ctx.Command.Context, ctx.Command.Cancel = context.WithCancel(context.Background())
	ctx.K8s.IOStreams = genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	ctx.K8s.RawConfig, err = configFlags.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return nil, err
	}
	ctx.K8s.RestConfig, err = configFlags.ToRESTConfig()
	if err != nil {
		return nil, err
	}
	return ctx, nil
}
