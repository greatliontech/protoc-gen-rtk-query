package module

import (
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
)

type moduleParams struct {
	Imports                map[string]string
	AddPbSuffix            bool
	WithMetadata           bool
	WithAIPStandardMethods bool
}

func parseParams(p pgs.Parameters) moduleParams {
	mp := moduleParams{
		Imports: map[string]string{},
	}

	if v, err := p.Bool("add_pb_suffix"); err == nil {
		mp.AddPbSuffix = v
	}

	for k, v := range p {
		if strings.HasPrefix(k, "M") {
			mp.Imports[k[1:]] = v
		}
	}

	if v, err := p.Bool("with_metadata"); err == nil {
		mp.WithMetadata = v
	}

	if v, err := p.Bool("with_aip_standard_methods"); err == nil {
		mp.WithAIPStandardMethods = v
	}

	return mp
}
