package managing

import "github.com/ktsymbal/seaports-catalogue/pkg/entities"

type MockService struct {
	Err   error
	Ports []entities.Port
}

func (m MockService) GetPort(id string) (entities.Port, error) {
	if m.Err != nil {
		return entities.Port{}, m.Err
	}
	return entities.MockPort, nil
}

func (m MockService) ListPorts(limit, offset uint) ([]entities.Port, error) {
	return m.Ports, m.Err
}

func (m MockService) CreatePort(entities.Port) error {
	return m.Err
}

func (m MockService) UpdatePort(port entities.Port) error {
	return m.Err
}

func (m MockService) CreateOrUpdatePorts(...entities.Port) error {
	return m.Err
}

func (m MockService) DeletePort(id string) error {
	return m.Err
}
