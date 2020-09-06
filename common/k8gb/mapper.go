package k8gb

import (
	"fmt"

	"k8s.io/client-go/rest"
)

//Mapper maps and validates from GslbRaw to GslbExt.
//Also extends by Ingress, CoreDns and k8gb pods
type Mapper struct {
	cluster string
	config *rest.Config
	raw GslbRaw
	ext GslbModel
}

func NewMapper(raw GslbRaw, config *rest.Config, cluster string) *Mapper {
	return &Mapper{
		cluster,
		config,
		raw,
		GslbModel{},
	}
}

func (m *Mapper) Map() GslbModel {
	return m.ext
}

func (m *Mapper) mapGslb() {
	m.ext.Kind = m.raw.kind()
	m.ext.Name = m.raw.name()
	m.ext.Namespace = m.raw.namespace()
	m.ext.GeoTag = m.raw.geoTag()
	m.ext.Host = m.host()
}


func (g *GslbRaw) geoTag() Stringr {
	return mandatoryStringr(g.Status.GeoTag,"empty geoTag")
}

func (g GslbRaw) kind() Stringr {
	return mandatoryStringr(g.Kind, "empty kind")
}

func (g GslbRaw) name() Stringr {
	return mandatoryStringr(g.Metadata.Name, "empty name")
}

func (g GslbRaw) namespace() Stringr {
	return mandatoryStringr(g.Metadata.Namespace, "empty namespace")
}

func (m *Mapper) host() Stringr {
	return mandatoryStringr(m.config.Host,"empty host")
}



// Basic validations, returning fields, taking care field is not empty etc..
func mandatoryStringr(str, err string) Stringr {
	s := Stringr{}
	s.Value = str
	if str == "" {
		s.IsValid = false
		s.Error = fmt.Errorf(err)
	}
	return s
}
