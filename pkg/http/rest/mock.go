package rest

import (
	"os"

	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
)

var mockPort = entities.Port{
	ID: "AEAJM", Name: "Ajman", City: "Ajman", Country: "United Arab Emirates", Coordinates: []float64{55.5136433,
		25.4052165}, Province: "Ajman", Timezone: "Asia/Dubai", Unlocs: []string{"AEAJM"}, Code: "52000"}

type MockParser struct{}

func (p MockParser) Parse(*os.File) []entities.Port {
	return []entities.Port{mockPort}
}

type MockRetriever struct{}

func (r MockRetriever) RetrievePort(id string) entities.Port {
	return mockPort
}

func (r MockRetriever) RetrieveAllPorts() []entities.Port {
	return []entities.Port{mockPort}
}
