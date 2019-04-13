package storage

import "fmt"

type ErrPortNotFound struct {
	ID string
}

func (e ErrPortNotFound) Error() string {
	return fmt.Sprintf("Port %s not found", e.ID)
}

type ErrPortAlreadyExists struct {
	ID string
}

func (e ErrPortAlreadyExists) Error() string {
	return fmt.Sprintf("Port %s already exists", e.ID)
}
