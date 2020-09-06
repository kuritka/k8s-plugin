package viewmodel

import (
	"fmt"

	"github.com/kuritka/plugin/common/k8gb"
	"github.com/kuritka/plugin/common/k8sctx"

	//"github.com/hashicorp/go-multierror"
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

func (vm *ViewModel) GetModel() (model map[string]map[string]k8gb.GslbModel, err error) {
	raw, err := vm.getRawGslbs()
	if err != nil {
		return
	}
	model, err = vm.mapToGslbModel(raw)
	return
}

func (vm *ViewModel) getRawGslbs() (raws map[string][]k8gb.GslbRaw, err error) {
	var unstructuredList *unstructured.UnstructuredList
	raws = make(map[string][]k8gb.GslbRaw,0)
	for _, ctx := range vm.k8s.RawConfig.Contexts {
		err = vm.k8s.SwitchContext(ctx.Cluster)
		if err != nil {
			return
		}
		unstructuredList,err = vm.k8s.DynamicConfig.Resource(k8gb.RuntimeClassGVR).List(metav1.ListOptions{})
		if err != nil {
			return
		}
		raws[ctx.Cluster] = getUnstructured(unstructuredList)
	}
	return
}

func (vm *ViewModel) mapToGslbModel(r map[string][]k8gb.GslbRaw) (m map[string]map[string]k8gb.GslbModel, err error) {
	m = make(map[string]map[string]k8gb.GslbModel,len(r))
	for cluster, gslbRaws := range r {
		m[cluster] = make(map[string]k8gb.GslbModel, len(gslbRaws))
		for _, raw := range gslbRaws {
			//TODO: inject validator
			m[cluster][raw.Metadata.Name] = k8gb.NewMapper(raw).Map()
		}
	}
	return
}


//maps unstructured data into GslbRaw structure. Any CRD change has to be reflected
//in GslbRaw or underlying structures
func getUnstructured(u *unstructured.UnstructuredList) (desc []k8gb.GslbRaw) {
	desc = make([]k8gb.GslbRaw, len(u.Items))
	for i, o := range u.Items {
		d := k8gb.GslbRaw{}
		d.Error = runtime.DefaultUnstructuredConverter.FromUnstructured(o.Object, &d)
		desc[i] = d
	}
	return
}


