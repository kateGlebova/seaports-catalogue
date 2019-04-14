package managing

import "github.com/kateGlebova/seaports-catalogue/pkg/entities"

type MockService struct {
	Err error
	Len int
}

func (m MockService) GetPort(id string) (entities.Port, error) {
	if m.Err != nil {
		return entities.Port{}, m.Err
	}
	return entities.MockPort, nil
}

func (m MockService) ListPorts(limit, offset uint) ([]entities.Port, error) {
	if m.Err != nil {
		return []entities.Port{}, m.Err
	}

	return entities.MockPorts(m.Len), nil
}

func (m MockService) CreatePort(entities.Port) error {
	return m.Err
}

func (m MockService) UpdatePort(id string) error {
	return m.Err
}

func (m MockService) CreateOrUpdatePorts(...entities.Port) error {
	return m.Err
}
