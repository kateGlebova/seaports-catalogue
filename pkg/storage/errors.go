package storage

type ErrPortNotFound struct{}

func (e ErrPortNotFound) Error() string {
	return "port not found"
}

type ErrPortAlreadyExists struct{}

func (e ErrPortAlreadyExists) Error() string {
	return "port already exists"
}
