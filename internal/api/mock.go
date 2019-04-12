package api

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

type MockRepo struct{}

func (r MockRepo) Get(id string) entities.Port {
	return mockPort
}

func (r MockRepo) GetAll() []entities.Port {
	return []entities.Port{mockPort}
}

func (r MockRepo) GetLimited(limit int) []entities.Port {
	return []entities.Port{mockPort}
}

func (r MockRepo) Save(...entities.Port) error {
	return nil
}
