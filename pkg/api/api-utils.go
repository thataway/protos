package api

import (
	_ "embed"
	"encoding/json"

	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
)

var (
	//go:embed announcer/api.swagger.json
	swaggerOfAnnouncer []byte

	//go:embed healthcheck/api.swagger.json
	swaggerOfHealthchecker []byte

	//go:embed ipruler/api.swagger.json
	swaggerOfIpruler []byte

	//go:embed ipvs/api.swagger.json
	swaggerOfIpvs []byte

	//go:embed route/api.swagger.json
	swaggerOfRoute []byte

	//go:embed tunnel/api.swagger.json
	swaggerOfTunnel []byte
)

var (
	api2swag = map[whatAPI][]byte{
		Announcer:     swaggerOfAnnouncer,
		Healthchecker: swaggerOfHealthchecker,
		Ipruler:       swaggerOfIpruler,
		Ipvs:          swaggerOfIpvs,
		Route:         swaggerOfRoute,
		Tunnel:        swaggerOfTunnel,
	}
	api2s = map[whatAPI]string{
		Announcer:     "Announcer",
		Healthchecker: "Healthchecker",
		Ipruler:       "Ipruler",
		Ipvs:          "Ipvs",
		Route:         "Route",
		Tunnel:        "Tunnel",
	}
)

var (
	//ErrSwaggerNotExist when document is not found
	ErrSwaggerNotExist = errors.New("swagger doc is no exist")
)

type whatAPI int

//nolint
const (
	Announcer whatAPI = iota + 1
	Healthchecker
	Ipruler
	Ipvs
	Route
	Tunnel
)

//LoadSwagger load swagger doc
func (a whatAPI) LoadSwagger() (*spec.Swagger, error) {
	const api = "LoadSwagger"
	raw, found := api2swag[a]
	if !found {
		return nil, ErrSwaggerNotExist
	}
	ret := new(spec.Swagger)
	err := json.Unmarshal(raw, ret)
	return ret, errors.Wrap(err, api)
}

//String ...
func (a whatAPI) String() string {
	s, ok := api2s[a]
	if !ok {
		return "?"
	}
	return s
}
