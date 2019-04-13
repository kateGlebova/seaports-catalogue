package adding

import (
	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
)

type Service interface {
	AddPorts(...entities.Port) error
}

type service struct {
	repository entities.PortRepository
}

func (s service) AddPorts(ports ...entities.Port) (err error) {
	for _, port := range ports {
		err = s.repository.SavePort(port)
		if err != nil {
			return
		}
	}
	return
}

func NewService() Service {
	return service{}
}
