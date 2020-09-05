package k8gb

import (
	"fmt"
)

//Mapper maps and validates from GslbRaw to GslbExt.
//Also extends by Ingress, CoreDns and k8gb pods
type Mapper struct {
	raw GslbRaw
	ext GslbExtended
}

func NewMapper(raw GslbRaw) *Mapper {
	return &Mapper{
		raw,
		GslbExtended{},
	}
}

func (m *Mapper) Map() GslbExtended {
	return m.ext
}

func (m *Mapper) mapGslb() {
	m.ext.Kind = m.raw.getKind()
	m.ext.Name = m.raw.getName()
	m.ext.Namespace = m.raw.getNamespace()
	//m.ext.GeoTag = m.raw.
	m.ext.Host = m.raw.getHost()
}

func (g GslbRaw) getKind() Stringr {
	return emptyStringValidator(g.Kind, "empty kind")
}

func (g GslbRaw) getName() Stringr {
	return emptyStringValidator(g.Metadata.Name, "empty name")
}

func (g GslbRaw) getNamespace() Stringr {
	return emptyStringValidator(g.Metadata.Namespace, "empty namespace")
}

func (g GslbRaw) getHost() Stringr {
	//u, err := url.ParseRequestURI(g.Metadata)
	return *new(Stringr)
}

func emptyStringValidator(str, err string) Stringr {
	s := Stringr{}
	s.Value = str
	if str == "" {
		s.IsValid = false
		s.Error = fmt.Errorf(err)
	}
	return s
}
