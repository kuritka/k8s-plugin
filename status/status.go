package status

import (
	"fmt"
	"strings"

	"github.com/kuritka/plugin/common/guard"
	"github.com/kuritka/plugin/common/k8gb"
	"github.com/kuritka/plugin/common/log"

	k8sctx2 "github.com/kuritka/plugin/common/k8sctx"
	"github.com/kyokomi/emoji"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

//StatusInfo command structure
type Info struct {
	options Options
}

//Options input vars for command
type Options struct {
	Namespace string
	Context   *k8sctx2.Context
}

var logger = log.Log

//New returns Status service implementation
func New(options Options) *Info {
	return &Info{
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

func printGslb(client dynamic.Interface) {
	rs := fmt.Sprintf("%s/%s", runtimeClassGVR.Group, runtimeClassGVR.Resource)
	res := client.Resource(runtimeClassGVR)
	list, err := res.List(metav1.ListOptions{})
	guard.FailOnError(err, "reading CRD")
	r := mapUnstructured(list)
	fmt.Print(r)
	logger.Info().Msgf("Printing %s.%s", rs, strings.Join([]string{"spec", "runtimeHandler"}, "."))
}

//Run runs the command implementation
func (s *Info) Run() error {
	emoji.Println(":unicorn:")
	logger.Info().Msgf(s.options.Namespace)
	for k := range s.options.Context.K8s.RawConfig.Clusters {
		logger.Info().Msgf(k)
	}
	cs, err := kubernetes.NewForConfig(s.options.Context.K8s.RestConfig)
	if err != nil {
		return err
	}

	//package api from k8gb import here...
	ing, err := cs.NetworkingV1beta1().Ingresses(s.options.Namespace).List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, n := range ing.Items {
		logger.Info().Msgf("%s %s", n.ClusterName, n.Name)
	}

	dc, err := dynamic.NewForConfig(s.options.Context.K8s.RestConfig)
	guard.FailOnError(err, "client")
	printGslb(dc)

	return nil
}

func (s *Info) String() string {
	return "Status"
}

//maps unstructured data into Desc structure. Any CRD change has to be reflected
//in Desc or underlying structures
func mapUnstructured(u *unstructured.UnstructuredList) (desc []k8gb.Desc) {
	desc = make([]k8gb.Desc, 2)
	for i, o := range u.Items {
		d := k8gb.Desc{}
		d.Error = runtime.DefaultUnstructuredConverter.FromUnstructured(o.Object, &d)
		desc[i] = d
	}
	return
}
