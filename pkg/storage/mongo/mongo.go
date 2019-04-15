package mongo

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
	"github.com/kateGlebova/seaports-catalogue/pkg/storage"
)

// Repository implements CRUD operations with MongoDB
type Repository struct {
	session *mgo.Session

	url        string
	db         string
	collection string
	err        error
}

// NewRepository instantiates new MongoDB repository
func NewRepository(url, db, collection string) (*Repository, error) {
	log.Print("Dialing MongoDB...")
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Repository{url: url, db: db, collection: collection, session: session}, nil
}

// Stop closes MongoDB session if one exists
func (r *Repository) Stop() error {
	if r.err != nil {
		return r.err
	}
	log.Print("Closing MongoDB session...")
	if r.session != nil {
		r.session.Close()
		return nil
	}
	log.Print("MongoDB session closed.")
	return nil
}

// GetPort finds port with passed id in MongoDB
func (r *Repository) GetPort(id string) (port entities.Port, err error) {
	sessionCopy := r.session.Copy()
	defer sessionCopy.Close()
	collection := r.session.DB(r.db).C(r.collection)

	err = collection.FindId(id).One(&port)
	if err == mgo.ErrNotFound {
		err = storage.ErrPortNotFound{}
	}

	return
}

// GetAllPorts gets from MongoDB a number of ports (defined by limit) starting from offset
func (r *Repository) GetAllPorts(limit, offset uint) (ports []entities.Port, err error) {
	sessionCopy := r.session.Copy()
	defer sessionCopy.Close()
	collection := r.session.DB(r.db).C(r.collection)

	if limit == 0 {
		return []entities.Port{}, storage.ErrNoLimit{}
	}

	err = collection.Find(nil).Sort("_id").Skip(int(offset)).Limit(int(limit)).All(&ports)
	if ports == nil {
		ports = []entities.Port{}
	}
	return
}

// CreatePort creates port if it doesn't exist in MongoDB
func (r *Repository) CreatePort(port entities.Port) (err error) {
	sessionCopy := r.session.Copy()
	defer sessionCopy.Close()
	collection := r.session.DB(r.db).C(r.collection)

	err = collection.FindId(port.ID).One(&port)
	if err != mgo.ErrNotFound {
		err = storage.ErrPortAlreadyExists{}
		return
	}

	return collection.Insert(port)
}

// UpdatePort updates existing port in MongoDB
func (r *Repository) UpdatePort(port entities.Port) (err error) {
	sessionCopy := r.session.Copy()
	defer sessionCopy.Close()
	collection := r.session.DB(r.db).C(r.collection)

	err = collection.UpdateId(port.ID, port)
	if err == mgo.ErrNotFound {
		err = storage.ErrPortNotFound{}
	}
	return
}

// CreateOrUpdatePorts creates ports that don't exist and updates ports that do
func (r *Repository) CreateOrUpdatePorts(ports ...entities.Port) (err error) {
	sessionCopy := r.session.Copy()
	defer sessionCopy.Close()
	bulk := r.session.DB(r.db).C(r.collection).Bulk()

	pairs := make([]interface{}, 0, len(ports)*2)
	for _, p := range ports {
		pairs = append(pairs, bson.M{"_id": p.ID}, p)
	}
	bulk.Upsert(pairs...)
	_, err = bulk.Run()
	return
}
