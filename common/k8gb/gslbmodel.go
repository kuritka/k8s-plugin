package k8gb

type GslbModel struct {
	Kind                       Stringr
	Cluster                    Stringr
	ApiVersion                 Stringr
	Name                       Stringr
	Host                       Stringr
	Namespace                  Stringr
	GeoTag                     Stringr
	Health                     map[string]RecordState
	DnsTtlSeconds              Intr
	SplitBrainThresholdSeconds Intr
	Type                       Stringr

	Ingress map[string]K8gbIngress
	Pods    map[string]K8gbPod
	CoreDns    map[string]CoreDnsPod
}

type RecordState struct {
	Status         Stringr
	HealthyRecords []string
}

type K8gbIngress struct {
	Status    Stringr
	Address   []Stringr
	Host      []Stringr
	Namespace string
}

type K8gbPod struct {
	Name      string
	Namespace string
}

type CoreDnsPod struct {
	Name      string
	Namespace string
}


type Intr struct {
	Value   int64
	Error   error
	IsValid bool
}

type Stringr struct {
	Value   string
	Error   error
	IsValid bool
}
