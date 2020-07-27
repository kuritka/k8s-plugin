package status

import (
	k8sctx "github.com/kuritka/plugin/common/k8s/k8s-context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/kuritka/plugin/common/log"
)

type Status struct {
	options Options
}

type Options struct {
	Namespace string
	Context *k8sctx.Context
}

var logger = log.Log

func New(options Options) *Status {
	return &Status{
		options,
	}
}

func (s *Status) Run() error {
	for k,_ := range s.options.Context.K8s.RawConfig.Clusters {
		logger.Info().Msgf(k)
	}
	clientset, err := kubernetes.NewForConfig(s.options.Context.K8s.RestConfig)
	if err != nil {
		return err
	}
	ns, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _,n := range ns.Items {
		logger.Info().Msgf("%s %s",n.ClusterName, n.Name)
	}
	return nil
}

func (s *Status) String() string {
	return "Status"
}