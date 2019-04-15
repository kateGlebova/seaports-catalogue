package mongo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kateGlebova/seaports-catalogue/pkg/storage"

	"github.com/kateGlebova/seaports-catalogue/pkg/entities"

	"github.com/globalsign/mgo"
)

var (
	repo *Repository
)

func TestMain(m *testing.M) {
	url, db := os.Getenv("MONGO_URL"), os.Getenv("MONGO_DB")
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	repo = &Repository{session: session, url: url, db: db, collection: "ports"}
	m.Run()
}

func TestRepository_GetPort(t *testing.T) {
	err := clearDB()
	if err != nil {
		t.Fatalf("error clearing DB: %v", err)
	}

	err = addPort(entities.MockPort)
	if err != nil {
		t.Fatalf("error clearing DB: %v", err)
	}

	testCases := []struct {
		name string
		id   string
		port entities.Port
		err  error
	}{
		{name: "port exists", id: entities.MockPort.ID, port: entities.MockPort},
		{name: "port doesn't exist", id: "id", port: entities.Port{}, err: storage.ErrPortNotFound{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			port, err := repo.GetPort(tc.id)
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.port, port)
		})
	}
}

func TestRepository_GetAllPorts(t *testing.T) {
	err := clearDB()
	if err != nil {
		t.Fatalf("error clearing DB: %v", err)
	}

	ports := entities.MockPorts(10)
	err = addPorts(ports)
	if err != nil {
		t.Fatalf("error adding ports to DB: %v", err)
	}

	testCases := []struct {
		name   string
		limit  uint
		offset uint
		ports  []entities.Port
		err    error
	}{
		{name: "no limit", limit: 0, err: storage.ErrNoLimit{}, ports: []entities.Port{}},
		{name: "offset 0, limit < len", offset: 0, limit: 3, ports: ports[:3]},
		{name: "offset 0, limit > len", offset: 0, limit: 12, ports: ports},
		{name: "0 < offset < len, limit < len", offset: 5, limit: 3, ports: ports[5:8]},
		{name: "0 < offset < len, limit > len", offset: 5, limit: 7, ports: ports[5:]},
		{name: "offset > len", offset: 11, limit: 7, ports: []entities.Port{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ports, err := repo.GetAllPorts(tc.limit, tc.offset)
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.ports, ports)
		})
	}
}

func TestRepository_GetAllPortsEmpty(t *testing.T) {
	err := clearDB()
	if err != nil {
		t.Fatalf("error clearing DB: %v", err)
	}

	testCases := []struct {
		name   string
		limit  uint
		offset uint
		ports  []entities.Port
		err    error
	}{
		{name: "no limit", limit: 0, err: storage.ErrNoLimit{}, ports: []entities.Port{}},
		{name: "offset 0, limit < len", offset: 0, limit: 3, ports: []entities.Port{}},
		{name: "offset 0, limit > len", offset: 0, limit: 12, ports: []entities.Port{}},
		{name: "0 < offset < len, limit < len", offset: 5, limit: 3, ports: []entities.Port{}},
		{name: "0 < offset < len, limit > len", offset: 5, limit: 7, ports: []entities.Port{}},
		{name: "offset > len", offset: 11, limit: 7, ports: []entities.Port{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ports, err := repo.GetAllPorts(tc.limit, tc.offset)
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.ports, ports)
		})
	}
}

func TestRepository_CreatePort(t *testing.T) {
	err := clearDB()
	if err != nil {
		t.Fatalf("error clearing DB: %v", err)
	}

	err = addPort(entities.MockPort)
	if err != nil {
		t.Fatalf("error adding port to DB: %v", err)
	}

	testCases := []struct {
		name string
		id   string
		err  error
	}{
		{name: "port exists", id: entities.MockPort.ID, err: storage.ErrPortAlreadyExists{}},
		{name: "port doesn't exist", id: "id"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			port := entities.MockPort
			port.ID = tc.id
			err := repo.CreatePort(port)
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
				var result entities.Port
				err = repo.session.DB(repo.db).C(repo.collection).FindId(tc.id).One(&result)
				assert.NoError(t, err)
				assert.Equal(t, port, result)
			}
		})
	}
}

func TestRepository_UpdatePort(t *testing.T) {
	err := clearDB()
	if err != nil {
		t.Fatalf("error clearing DB: %v", err)
	}

	err = addPort(entities.MockPort)
	if err != nil {
		t.Fatalf("error adding port to DB: %v", err)
	}

	port := entities.Port{
		Name:        "Name",
		City:        "City",
		Country:     "Country",
		Alias:       make([]string, 0),
		Regions:     make([]string, 0),
		Coordinates: []float64{25.4052165, 55.5136433},
		Province:    "Province",
		Timezone:    "Europe/Kyiv",
		Unlocs:      []string{"AEAJL"},
		Code:        "52005",
	}

	testCases := []struct {
		name string
		id   string
		err  error
	}{
		{name: "port exists", id: entities.MockPort.ID},
		{name: "port doesn't exist", id: "id", err: storage.ErrPortNotFound{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			instance := port
			instance.ID = tc.id
			err := repo.UpdatePort(instance)
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
				var result entities.Port
				err = repo.session.DB(repo.db).C(repo.collection).FindId(tc.id).One(&result)
				assert.NoError(t, err)
				assert.Equal(t, instance, result)
			}
		})
	}
}

func TestRepository_CreateOrUpdatePorts(t *testing.T) {
	err := clearDB()
	if err != nil {
		t.Fatalf("error clearing DB: %v", err)
	}

	err = addPort(entities.MockPort)
	if err != nil {
		t.Fatalf("error adding port to DB: %v", err)
	}

	port := entities.Port{
		Name:        "Name",
		City:        "City",
		Country:     "Country",
		Alias:       make([]string, 0),
		Regions:     make([]string, 0),
		Coordinates: []float64{25.4052165, 55.5136433},
		Province:    "Province",
		Timezone:    "Europe/Kyiv",
		Unlocs:      []string{"AEAJL"},
		Code:        "52005",
	}

	testCases := []struct {
		name string
		id   string
	}{
		{name: "port exists", id: entities.MockPort.ID},
		{name: "port doesn't exist", id: "id"},
	}

	ports := make([]entities.Port, 0, len(testCases))
	for _, tc := range testCases {
		instance := port
		instance.ID = tc.id
		ports = append(ports, instance)
	}

	err = repo.CreateOrUpdatePorts(ports...)
	assert.NoError(t, err)

	for _, p := range ports {
		var result entities.Port
		err = repo.session.DB(repo.db).C(repo.collection).FindId(p.ID).One(&result)
		assert.NoError(t, err)
		assert.Equal(t, p, result)
	}
}

func clearDB() error {
	return repo.session.DB(repo.db).DropDatabase()
}

func addPorts(ports []entities.Port) error {
	iports := make([]interface{}, len(ports))
	for i, v := range ports {
		iports[i] = v
	}
	return repo.session.DB(repo.db).C(repo.collection).Insert(iports...)
}

func addPort(port entities.Port) error {
	return repo.session.DB(repo.db).C(repo.collection).Insert(port)
}
