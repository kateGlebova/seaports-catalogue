package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
	"github.com/kateGlebova/seaports-catalogue/pkg/storage"
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

func (r *Repository) GetPort(id string) (port entities.Port, err error) {
	sessionCopy := r.session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(r.db).C(r.collection)

	err = collection.FindId(id).One(&port)
	if err == mgo.ErrNotFound {
		err = storage.ErrPortNotFound{}
	}

	return
}

func (r *Repository) GetAllPorts(limit, offset uint) (ports []entities.Port, err error) {
	sessionCopy := r.session.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(r.db).C(r.collection)

	if limit == 0 {
		return []entities.Port{}, storage.ErrNoLimit{}
	}

	err = collection.Find(nil).Skip(int(offset)).Limit(int(limit)).All(&ports)
	if ports == nil {
		ports = []entities.Port{}
	}
	return
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
