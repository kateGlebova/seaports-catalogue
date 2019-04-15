package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
)

type Repository struct {
	session *mgo.Session

	url        string
	db         string
	collection string
	err        error
}

func NewRepository(url, db, collection string) *Repository {
	return &Repository{url: url, db: db, collection: collection}
}

func (r *Repository) GetPort(id string) (entities.Port, error) {
	return entities.Port{}, nil
}

func (r *Repository) GetAllPorts(limit, offset uint) ([]entities.Port, error) {
	return []entities.Port{}, nil
}

func (r *Repository) CreatePort(port entities.Port) error {
	return nil
}

func (r *Repository) UpdatePort(port entities.Port) error {
	return nil
}

func (r *Repository) CreateOrUpdatePorts(ports ...entities.Port) error {
	return nil
}
