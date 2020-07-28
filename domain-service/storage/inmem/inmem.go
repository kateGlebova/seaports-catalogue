package inmem

import (
	"github.com/ktsymbal/seaports-catalogue/domain-service/storage"
	"sync"

	"github.com/ktsymbal/seaports-catalogue/pkg/entities"
)

type Repository struct {
	ports map[string]entities.Port
	mutex sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{ports: make(map[string]entities.Port)}
}

func (r *Repository) GetPort(id string) (entities.Port, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	port, ok := r.ports[id]
	if !ok {
		return port, storage.ErrPortNotFound{}
	}
	return port, nil
}

func (r *Repository) GetAllPorts(limit, offset uint) ([]entities.Port, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	ports := make([]entities.Port, 0, len(r.ports))
	for _, port := range r.ports {
		ports = append(ports, port)
	}
	return ports, nil
}

func (r *Repository) CreatePort(port entities.Port) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.ports[port.ID]; ok {
		return storage.ErrPortAlreadyExists{}
	}

	r.ports[port.ID] = port
	return nil
}

func (r *Repository) UpdatePort(port entities.Port) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.ports[port.ID]; !ok {
		return storage.ErrPortNotFound{}
	}

	r.ports[port.ID] = port
	return nil
}

func (r *Repository) CreateOrUpdatePorts(ports ...entities.Port) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, port := range ports {
		r.ports[port.ID] = port
	}
	return nil
}
