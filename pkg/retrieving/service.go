package retrieving

import (
	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
)

type Service interface {
	RetrievePort(name string) entities.Port
	RetrieveAllPorts() []entities.Port
}

type service struct {
	// grpc client connection
}

func (s service) RetrievePort(id string) entities.Port {
	// grpc request to retrieve the port
	return entities.Port{}
}

func (s service) RetrieveAllPorts() []entities.Port {
	// grpc request to retrieve all ports
	return []entities.Port{}
}

func NewService() Service {
	return service{}
}
