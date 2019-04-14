package rest

import (
	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
)

var mockPort = entities.Port{
	ID: "AEAJM", Name: "Ajman", City: "Ajman", Country: "United Arab Emirates", Coordinates: []float64{55.5136433,
		25.4052165}, Province: "Ajman", Timezone: "Asia/Dubai", Unlocs: []string{"AEAJM"}, Code: "52000"}

func mockPorts() []entities.Port {
	length := 5
	ports := make([]entities.Port, 0, length)
	for i := 0; i < length; i++ {
		ports = append(ports, mockPort)
	}
	return ports
}

type MockManager struct {
	err   error
	empty bool
}

func (m MockManager) GetPort(id string) (entities.Port, error) {
	if m.err != nil {
		return entities.Port{}, m.err
	}
	return mockPort, nil
}

func (m MockManager) ListPorts(limit, offset uint) ([]entities.Port, error) {
	if m.err != nil {
		return []entities.Port{}, m.err
	}

	if m.empty {
		return []entities.Port{}, nil
	}

	return mockPorts(), nil
}

func (m MockManager) CreatePort(entities.Port) error {
	return m.err
}

func (m MockManager) UpdatePort(id string) error {
	return m.err
}

func (m MockManager) CreateOrUpdatePorts(...entities.Port) error {
	return m.err
}

type testError struct{}

func (err testError) Error() string {
	return "test error"
}

type errReader struct{}

func (r errReader) Read([]byte) (int, error) {
	return 0, testError{}
}
