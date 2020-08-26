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
	ctxBackup  string
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

func (k *K8s) SwitchContext(ctx string) (err error) {
	override := &clientcmd.ConfigOverrides{CurrentContext: ctx}
	clientConfig := clientcmd.NewNonInteractiveClientConfig(k.RawConfig, override.CurrentContext,
		override, &clientcmd.ClientConfigLoadingRules{})
	k.RawConfig, err = clientConfig.RawConfig()
	if err != nil {
		return
	}
	k.ClientConfig, err = clientConfig.ClientConfig()
	if err != nil {
		return
	}
	k.DynamicConfig, err = dynamic.NewForConfig(k.ClientConfig)
	return
}

//TearDown would be primarily deferred
func (k *K8s) TearDown() error {
	return k.SwitchContext(k.ctxBackup)
}
