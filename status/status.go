package status

import (
	"fmt"
	"github.com/kuritka/plugin/common/guard"
	"github.com/kuritka/plugin/common/k8gb"
	k8sctx2 "github.com/kuritka/plugin/common/k8sctx"
	"github.com/kuritka/plugin/status/internal/printer"
	"k8s.io/client-go/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
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
func (s *Info) Run() (err error) {

	cs, err := kubernetes.NewForConfig(s.options.Context.K8s.ClientConfig)
	if err != nil {
		return err
	}
	ing,err :=cs.NetworkingV1beta1().Ingresses(s.options.Namespace).List(metav1.ListOptions{})


	p := printer.DefaultPrettyPrinter()
	gslb := s.getGslb()
	guard.HandleError(p.Title(fmt.Sprintf("context: %s (%s)",s.options.Context.K8s.RawConfig.CurrentContext,s.options.Context.K8s.ClientConfig.Host)))
	for _,g := range gslb {
		guard.HandleError(p.Subtitle(fmt.Sprintf("%s %s:%s in namespace: %s",g.Metadata.Na§me, g.ApiVersion,g.Kind, g.Metadata.Namespace)))
		guard.HandleError(p.Property("Type",g.Type))
		guard.HandleError(p.Property("GeoTag",g.GeoTag))
		guard.HandleError(p.Property("DnsTTL",intToSec(g.DnsTtlSeconds)))
		guard.HandleError(p.Property("SplitBrain",intToSec(g.SplitBrainThresholdSeconds)))
		guard.HandleError(p.PropertyMap("Health", g.ServiceHealth))
		for k,h := range g.HealthyRecords {
			guard.HandleError(p.PropertySlice(k, h))
		}
		for _, item := range ing.Items {
			guard.HandleError(p.Property(item.Name,item.Status.LoadBalancer.Ingress[0].IP))
		}
		p.NewLine()
	}
	return
}


func intToSec(v int64) string{
	if v == 0 {
		return fmt.Sprintf(" - ")
	}
	return fmt.Sprintf("%vs",v)
}

//func print
func (s *Info) String() string {
	return "Status"
}

func (s *Info) getGslb() (gslb []k8gb.GslbRaw){
	res := s.options.Context.K8s.DynamicConfig.Resource(k8gb.RuntimeClassGVR)
	list, err := res.List(metav1.ListOptions{})
	guard.FailOnError(err, "reading CRD")
	return mapUnstructured(list)
}

//maps unstructured data into GslbRaw structure. Any CRD change has to be reflected
//in GslbRaw or underlying structures
func mapUnstructured(u *unstructured.UnstructuredList) (desc []k8gb.GslbRaw) {
	desc = make([]k8gb.GslbRaw, len(u.Items))
	for i, o := range u.Items {
		d := k8gb.GslbRaw{}
		d.Error = runtime.DefaultUnstructuredConverter.FromUnstructured(o.Object, &d)
		desc[i] = d
	}
	return
}
