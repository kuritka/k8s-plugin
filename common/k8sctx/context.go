package k8sctx

import (
	"context"

	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd/api"
)

//K8s contains k8s context
type K8s struct {
	ResultingContext     *api.Context
	ResultingContextName string

	DynamicConfig  dynamic.Interface
	ClientConfig   *rest.Config
	RawConfig      api.Config
	ListNamespaces bool
	genericclioptions.IOStreams
	kubeConfig string
}

//Command contains command
type Command struct {
	Args    []string
	Context context.Context
	Cancel  context.CancelFunc
}

//Context contains fill command context
type Context struct {
	K8s     *K8s
	Command *Command
}

func (k K8s) SwitchContext(ctx string) (err error) {
	cfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: k.kubeConfig},
		&clientcmd.ConfigOverrides{CurrentContext: ctx})
	k.RawConfig, err = cfg.RawConfig()
	k.ClientConfig, err = cfg.ClientConfig()
	return
}
