package k8gb

type Desc struct {
	Name       string
	Kind       string
	Cluster    string
	ApiVersion string
	Error      error
	Metadata
	Status
	Spec
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
	Ingress
	Strategy
}

type Strategy struct {
	DnsTtlSeconds              int64
	SplitBrainThresholdSeconds int64
	Type                       string
}

type Ingress struct {
	Rules []Rule
}

type Rule struct {
	Host string
	Http
}

type Http struct {
	Paths []Path
}

type Path struct {
	Backend
	Path string
}

type Backend struct {
	ServiceName string
	ServicePort int64
}
