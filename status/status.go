package status

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/kuritka/plugin/common/guard"
	"github.com/kuritka/plugin/common/k8gb"
	k8sctx2 "github.com/kuritka/plugin/common/k8sctx"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
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

//New returns Status service implementation
func New(options Options) *Info {
	return &Info{
		options,
	}
}

//Run runs the command implementation
func (s *Info) Run() error {
	printGslb(s.options.Context.K8s.DynamicConfig)
	fmt.Println(s.options.Context.K8s.RawConfig.CurrentContext)
	e := s.options.Context.K8s.SwitchContext("kind-test-gslb2")
	guard.FailOnError(e, "")
	fmt.Println(s.options.Context.K8s.RawConfig.CurrentContext)
	printGslb(s.options.Context.K8s.DynamicConfig)
	return nil
}

//func print

func (s *Info) String() string {
	return "Status"
}

func printGslb(client dynamic.Interface) {
	res := client.Resource(k8gb.RuntimeClassGVR)
	list, err := res.List(metav1.ListOptions{})
	guard.FailOnError(err, "reading CRD")
	r := mapUnstructured(list)
	for _, ru := range r {
		s := fmt.Sprintf("%v %s %s", emoji.Unicorn, ru.Metadata.Name, ru.Status.GeoTag)
		fmt.Println(s)
	}
}

//maps unstructured data into Desc structure. Any CRD change has to be reflected
//in Desc or underlying structures
func mapUnstructured(u *unstructured.UnstructuredList) (desc []k8gb.Desc) {
	desc = make([]k8gb.Desc, len(u.Items))
	for i, o := range u.Items {
		d := k8gb.Desc{}
		d.Error = runtime.DefaultUnstructuredConverter.FromUnstructured(o.Object, &d)
		desc[i] = d
	}
	return
}
