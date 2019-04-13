package parsing

import (
	"os"

	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
)

type Service interface {
	Parse(*os.File) []entities.Port
}

type service struct{}

func (s service) Parse(*os.File) []entities.Port {
	return []entities.Port{}
}

func NewService() Service {
	return service{}
}
