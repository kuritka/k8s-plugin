package k8gb

import "k8s.io/apimachinery/pkg/runtime/schema"

var RuntimeClassGVR = schema.GroupVersionResource{
	Group:    "k8gb.absa.oss",
	Version:  "v1beta1",
	Resource: "gslbs",
}
