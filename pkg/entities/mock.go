package entities

import (
	uuid "github.com/satori/go.uuid"
)

var MockPort = Port{
	ID:          "AEAJM",
	Name:        "Ajman",
	City:        "Ajman",
	Country:     "United Arab Emirates",
	Alias:       make([]string, 0),
	Regions:     make([]string, 0),
	Coordinates: []float64{55.5136433, 25.4052165},
	Province:    "Ajman",
	Timezone:    "Asia/Dubai",
	Unlocs:      []string{"AEAJM"},
	Code:        "52000",
}

func MockPorts(len int) []Port {
	ports := make([]Port, 0, len)
	for i := 0; i < len; i++ {
		port := MockPort
		port.ID = uuid.NewV4().String()
		ports = append(ports, port)
	}
	return ports
}
