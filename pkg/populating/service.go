package populating

import (
	"log"
	"os"

	"github.com/kateGlebova/seaports-catalogue/pkg/managing"
)

type Service interface {
	Populate() error
}

func NewService(fileName string, manager managing.Service) Service {
	return &jsonService{fileName: fileName, manager: manager}
}

type jsonService struct {
	fileName string
	file     *os.File
	manager  managing.Service
}

func (s *jsonService) Populate() error {
	return nil
}

func (s *jsonService) Run() {
	file, err := os.Open(s.fileName)
	if err != nil {
		log.Printf("error opening file '%s': %v", s.fileName, err)
		return
	}
	s.file = file

	defer func() { s.file.Close(); s.file = nil }()

	if err := s.Populate(); err != nil {
		log.Printf("error populating the database: %v", err)
	}
}

func (s *jsonService) Stop() error {
	log.Printf("Closing '%s' file...", s.fileName)
	if s.file != nil {
		return s.file.Close()
	}
	return nil
}
