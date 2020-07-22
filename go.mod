module github.com/kuritka/plugin

go 1.14

require (
	github.com/rs/zerolog v1.18.0
	github.com/spf13/cobra v0.0.6
    k8s.io/api v0.17.2
    k8s.io/apimachinery v0.17.2
    k8s.io/client-go v0.17.2
)

replace (
	k8s.io/api => k8s.io/api v0.17.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.2
	k8s.io/client-go => k8s.io/client-go v0.17.2
)
