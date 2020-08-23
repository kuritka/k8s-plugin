package status

import (
	"fmt"
	"github.com/kuritka/plugin/common/guard"
	k8sctx2 "github.com/kuritka/plugin/common/k8sctx"
	"github.com/kyokomi/emoji"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"strings"

	"github.com/kuritka/plugin/common/log"
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

type Metadata struct {
	Name        string
	Namespace   string
	Annotations map[string]string
}

type Status struct {
	GeoTag         string
	HealthyRecords map[string][]string
	ServiceHealth  map[string]string
}

type Spec struct {
	//Ingress map[string][]string
	Strategy
}

type Strategy struct {
	DnsTtlSeconds              string
	SplitBrainThresholdSeconds string
	Type                       string
}

type description struct {
	Name       string
	Kind       string
	Cluster    string
	ApiVersion string
	Metadata
	Status
	Spec
}

//gets
func mapUnstructured(un *unstructured.UnstructuredList) (desc []description) {
	desc = make([]description, 2)
	for i, u := range un.Items {
		d := description{}
		//d.Name = u.GetName()
		//d.Kind = fmt.Sprintf("%v", u.Object["kind"])
		//d.ApiVersion = fmt.Sprintf("%v", u.Object["apiVersion"])
		//d.Value = fmt.Sprintf("%v", interface(u.Object["metadata"])["name"])

		err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &d)
		guard.FailOnError(err, "")
		//	err := runtime.DefaultUnstructuredConverter.FromUnstructured(map[string](u.Object["metadata"]), &d)
		//	guard.FailOnError(err, "")
		desc[i] = d
	}
	return
}
