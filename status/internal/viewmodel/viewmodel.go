package viewmodel

import (
	"fmt"
	"github.com/kuritka/plugin/common/k8gb"
	"github.com/kuritka/plugin/common/k8sctx"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

type ViewModel struct {
	k8s *k8sctx.K8s
}

func NewViewModel(k8s *k8sctx.K8s) (vm *ViewModel, err error) {
	if k8s == nil {
		err = fmt.Errorf("k8s is nil")
		return
	}
	vm = new(ViewModel)
	vm.k8s = k8s
	return
}

func (vm *ViewModel) GetRawGslbs() (graw map[string][]k8gb.GslbRaw, err error) {
	var unstructuredList *unstructured.UnstructuredList
	graw = make(map[string][]k8gb.GslbRaw,0)
	for _, ctx := range vm.k8s.RawConfig.Contexts {
		err = vm.k8s.SwitchContext(ctx.Cluster)
		if err != nil {
			return
		}
		unstructuredList,err = vm.k8s.DynamicConfig.Resource(k8gb.RuntimeClassGVR).List(metav1.ListOptions{})
		if err != nil {
			return
		}
		graw[ctx.Cluster] = mapUnstructured(unstructuredList)
	}
	return
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


func (vm *ViewModel) GetModel() k8gb.GslbExtended {
	return k8gb.GslbExtended{}
}