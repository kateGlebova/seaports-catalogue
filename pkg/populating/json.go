package populating

import (
	"log"
	"os"
	"sync"

	"github.com/kateGlebova/seaports-catalogue/pkg/entities"

	"github.com/bcicen/jstream"

	"github.com/kateGlebova/seaports-catalogue/pkg/managing"
)

type jsonService struct {
	fileName string
	file     *os.File
	manager  managing.Service
}

// Run opens JSON file to parse and populates the database
func (s *jsonService) Run() {
	file, err := os.Open(s.fileName)
	if err != nil {
		log.Printf("error opening file '%s': %v", s.fileName, err)
		return
	}
	s.file = file

	defer func() { s.file.Close(); s.file = nil }()

	log.Printf("Populating from %s...", s.fileName)
	for err := s.Populate(); err != nil; err = s.Populate() {
		log.Printf("error populating the database: %v", err)
	}
	log.Print("Populating finished")
}

// Stop closes JSON file if it's open
func (s *jsonService) Stop() error {
	log.Printf("Closing '%s' file...", s.fileName)
	if s.file != nil {
		return s.file.Close()
	}
	return nil
}

// Populate parses JSON file and creates corresponding ports in the database
func (s *jsonService) Populate() error {
	wg := sync.WaitGroup{}
	err := make(chan error)
	quit := make(chan int)
	dec := jstream.NewDecoder(s.file, 1).EmitKV()

	for entry := range dec.Stream() {
		wg.Add(1)
		go func(port jstream.KV, errChan chan error) {
			mapPort := jsonPort(port.Value.(map[string]interface{}))
			p := mapPort.ToDomainPort(port.Key)
			err := s.manager.CreateOrUpdatePorts(p)
			if err != nil {
				errChan <- err
			}
			wg.Done()
		}(entry.Value.(jstream.KV), err)
	}

	go func(quit chan int) {
		wg.Wait()
		quit <- 0
	}(quit)

	select {
	case e := <-err:
		return e
	case <-quit:
		return nil
	}
}

type jsonPort map[string]interface{}

func (p jsonPort) ToDomainPort(id string) entities.Port {
	return entities.Port{
		ID:          id,
		Name:        p.getStringField("name"),
		City:        p.getStringField("city"),
		Country:     p.getStringField("country"),
		Alias:       p.getStringSliceField("alias"),
		Regions:     p.getStringSliceField("regions"),
		Coordinates: p.getFloatSliceField("coordinates"),
		Province:    p.getStringField("province"),
		Timezone:    p.getStringField("timezone"),
		Unlocs:      p.getStringSliceField("unlocs"),
		Code:        p.getStringField("code"),
	}
}

func (p jsonPort) getStringField(field string) (value string) {
	if p[field] != nil {
		value = p[field].(string)
	}
	return
}

func (p jsonPort) getStringSliceField(field string) []string {
	if p[field] == nil {
		return nil
	}
	values := make([]string, len(p[field].([]interface{})))
	for _, v := range p[field].([]interface{}) {
		values = append(values, v.(string))
	}
	return values
}

func (p jsonPort) getFloatSliceField(field string) []float64 {
	if p[field] == nil {
		return nil
	}
	values := make([]float64, len(p[field].([]interface{})))
	for _, v := range p[field].([]interface{}) {
		values = append(values, v.(float64))
	}
	return values
}
