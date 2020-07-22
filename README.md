#Plugin


## running the project

```shell script
go mod download
```


### Troubleshoot
If you are unable to build project because of dependency issues run following
```shell script
go clean -modcache
```
ensure you have go 1.14 within go.mod and file looks like following:

```go
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
``` 
than run 

```shell script
go mod download
```

### todo
update `k8s.io/api` version 


### refs
https://github.com/vladimirvivien/k8s-client-examples
https://medium.com/programming-kubernetes/building-stuff-with-the-kubernetes-api-part-4-using-go-b1d0e3c1c899
https://soggy.space/namespaced-crds-dynamic-client/