package populating

import (
	"github.com/kateGlebova/seaports-catalogue/pkg/managing"
)

type Service interface {
	Populate() error
}

func NewService(fileName string, manager managing.Service) Service {
	return &jsonService{fileName: fileName, manager: manager}
}
