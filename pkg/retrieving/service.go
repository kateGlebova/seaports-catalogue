package retrieving

import (
	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
)

type Service interface {
	RetrievePort(name string) entities.Port
	RetrieveAllPorts() []entities.Port
}

type service struct {
	repository entities.PortRepository
}

func (s service) RetrievePort(id string) entities.Port {
	return s.repository.GetPort(id)
}

func (s service) RetrieveAllPorts() []entities.Port {
	return s.repository.GetAllPorts()
}

func NewService() Service {
	return service{}
}
