package adding

import (
	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
)

type Service interface {
	AddPorts(...entities.Port) error
}

type service struct {
	// grpc client connection
}

func (s service) AddPorts(ports ...entities.Port) (err error) {
	// grpc request to add port
	return nil
}

func NewService() Service {
	return service{}
}
