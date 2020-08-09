package status

import (
	k8sctx2 "github.com/kuritka/plugin/common/k8sctx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/kuritka/plugin/common/log"
)

//Status command structure
type Status struct {
	options Options
}

//Options input vars for command
type Options struct {
	Namespace string
	Context   *k8sctx2.Context
}

var logger = log.Log

//New returns Status service implementation
func New(options Options) *Status {
	return &Status{
		options,
	}
}

//Run runs the command implementation
func (s *Status) Run() error {
	logger.Info().Msgf(s.options.Namespace)
	for k := range s.options.Context.K8s.RawConfig.Clusters {
		logger.Info().Msgf(k)
	}
	cs, err := kubernetes.NewForConfig(s.options.Context.K8s.RestConfig)
	if err != nil {
		return err
	}
	//ns, err := cs.CoreV1().Namespaces().List(metav1.ListOptions{})
	//if err != nil {
	//	return err
	//}
	//
	//package api from k8gb import here...
	ing, err := cs.NetworkingV1beta1().Ingresses(s.options.Namespace).List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, n := range ing.Items {
		logger.Info().Msgf("%s %s", n.ClusterName, n.Name)
	}
	return nil
}

func (s *Status) String() string {
	return "Status"
}
