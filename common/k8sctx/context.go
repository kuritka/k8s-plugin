package k8sctx

import (
	"context"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd/api"
)

//K8s contains k8s context
type K8s struct {
	ResultingContext     *api.Context
	ResultingContextName string

	RestConfig		*rest.Config
	RawConfig		api.Config
	ListNamespaces bool
	genericclioptions.IOStreams
}

//Command contains command
type Command struct {
	Args           []string
	Context context.Context
	Cancel context.CancelFunc
}

//Context contains fill command context
type Context struct{
	K8s *K8s
	Command *Command
}
