package k8gb

//Mapper maps and validates from GslbRaw to GslbExt.
//Also extends by Ingress, CoreDns and k8gb pods
type Mapper struct {
	raw GslbRaw
}

func NewMapper(raw GslbRaw) *Mapper {
	return &Mapper{
		raw,
	}
}

func (m *Mapper) Map() GslbExtended{
	return GslbExtended{}
}
