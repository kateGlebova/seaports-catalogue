package proto

import (
	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
	"github.com/kateGlebova/seaports-catalogue/pkg/storage"
	"google.golang.org/grpc/codes"
)

var (
	errorMapping = map[string]codes.Code{
		storage.ErrPortNotFound{}.Error():      codes.NotFound,
		storage.ErrPortAlreadyExists{}.Error(): codes.AlreadyExists,
	}
)

func DomainToProtoPort(port entities.Port) *Port {
	return &Port{
		Id:          port.ID,
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Coordinates: port.Coordinates,
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}
}

func ProtoToDomainPort(port *Port) entities.Port {
	return entities.Port{
		ID:          port.Id,
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Coordinates: port.Coordinates,
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}
}
