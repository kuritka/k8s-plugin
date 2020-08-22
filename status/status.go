package status

import (
	"fmt"
	"github.com/kuritka/plugin/common/guard"
	k8sctx2 "github.com/kuritka/plugin/common/k8sctx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"strings"

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

var (
	runtimeClassGVR = schema.GroupVersionResource{
		Group:    "k8gb.absa.oss",
		Version:  "v1beta1",
		Resource: "gslbs",
	}
)

func x(client dynamic.Interface) {
	rs := fmt.Sprintf("%s/%s", runtimeClassGVR.Group, runtimeClassGVR.Resource)
	res := client.Resource(runtimeClassGVR)
	list, err := res.List(metav1.ListOptions{})
	guard.FailOnError(err, "reading CRD")
	fmt.Print(list)
	logger.Info().Msgf("Printing %s.%s", rs, strings.Join([]string{"spec", "runtimeHandler"}, "."))

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

	client, err := dynamic.NewForConfig(s.options.Context.K8s.RestConfig)
	guard.FailOnError(err, "client")
	x(client)
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
