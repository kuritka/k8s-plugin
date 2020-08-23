module github.com/kuritka/plugin

go 1.15

require (
	github.com/kyokomi/emoji v2.2.4+incompatible
	github.com/rs/zerolog v1.18.0
	github.com/spf13/cobra v0.0.6
	k8s.io/api v0.17.2
	k8s.io/apiextensions-apiserver v0.17.2 // https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1?tab=doc
	k8s.io/apimachinery v0.17.2
	k8s.io/cli-runtime v0.17.2
	k8s.io/client-go v0.17.2
)

replace (
	k8s.io/api => k8s.io/api v0.17.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.2
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.17.2
	k8s.io/client-go => k8s.io/client-go v0.17.2
)
