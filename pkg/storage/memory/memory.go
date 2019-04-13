package memory

import (
	"sync"

	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
	"github.com/kateGlebova/seaports-catalogue/pkg/storage"
)

type Repository struct {
	ports map[string]entities.Port
	mutex sync.RWMutex
}

func (r Repository) GetPort(id string) (entities.Port, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	port, ok := r.ports[id]
	if !ok {
		return port, storage.ErrPortNotFound{ID: id}
	}
	return port, nil
}

func (r Repository) GetAllPorts() []entities.Port {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	ports := make([]entities.Port, 0, len(r.ports))
	for _, port := range r.ports {
		ports = append(ports, port)
	}
	return ports
}

func (r Repository) CreatePort(port entities.Port) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if port, ok := r.ports[port.ID]; ok {
		return storage.ErrPortAlreadyExists{ID: port.ID}
	}

	r.ports[port.ID] = port
	return nil
}

func (r Repository) UpdatePort(port entities.Port) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if port, ok := r.ports[port.ID]; !ok {
		return storage.ErrPortNotFound{ID: port.ID}
	}

	r.ports[port.ID] = port
	return nil
}

func (r Repository) CreateOrUpdate(port entities.Port) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.ports[port.ID] = port
	return nil
}
