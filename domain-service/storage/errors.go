package storage

import (
	"google.golang.org/grpc/codes"
)

type ErrPortNotFound struct{}

func (e ErrPortNotFound) Error() string {
	return "port not found"
}

type ErrPortAlreadyExists struct{}

func (e ErrPortAlreadyExists) Error() string {
	return "port already exists"
}

type ErrNoLimit struct{}

func (e ErrNoLimit) Error() string {
	return "no limit specified"
}

var (
	GRPCErrorMapping = map[string]codes.Code{
		ErrPortNotFound{}.Error():      codes.NotFound,
		ErrPortAlreadyExists{}.Error(): codes.AlreadyExists,
		ErrNoLimit{}.Error():           codes.InvalidArgument,
	}
)
