package proto

import (
	"context"
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
)

const (
	DefaultLimit  = 50
	DefaultOffset = 0
)

type RepositoryService struct {
	portRepo entities.PortRepository
}

func NewRepositoryService(portRepo entities.PortRepository) *RepositoryService {
	return &RepositoryService{portRepo: portRepo}
}

func (s RepositoryService) ListPorts(r *ListRequest, stream Repository_ListPortsServer) error {
	limit, offset := uint(r.Limit), uint(r.Offset)
	if limit == 0 {
		limit = DefaultLimit
	}
	if offset == 0 {
		offset = DefaultOffset
	}

	for _, port := range s.portRepo.GetAllPorts(limit, offset) {
		if err := stream.Send(DomainToProtoPort(port)); err != nil {
			return err
		}
	}
	return nil
}

func (s RepositoryService) GetPort(_ context.Context, port *Port) (*Port, error) {
	p, err := s.portRepo.GetPort(port.Id)
	if err != nil {
		return port, gRPCError(err)
	}
	return DomainToProtoPort(p), err
}

func (s RepositoryService) CreatePort(_ context.Context, port *Port) (*Empty, error) {
	p := ProtoToDomainPort(port)
	err := s.portRepo.CreatePort(p)
	if err != nil {
		return &Empty{}, gRPCError(err)
	}
	return &Empty{}, nil
}

func (s RepositoryService) UpdatePort(_ context.Context, port *Port) (*Empty, error) {
	p := ProtoToDomainPort(port)
	err := s.portRepo.UpdatePort(p)
	if err != nil {
		return &Empty{}, gRPCError(err)
	}
	return &Empty{}, nil
}

func (s RepositoryService) CreateOrUpdatePorts(stream Repository_CreateOrUpdatePortsServer) error {
	for {
		port, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&Empty{})
		}
		if err != nil {
			return err
		}
		err = s.portRepo.CreateOrUpdate(ProtoToDomainPort(port))
		if err != nil {
			return gRPCError(err)
		}
	}
}

func gRPCError(err error) error {
	errMessage := err.Error()
	code, ok := errorMapping[errMessage]
	if !ok {
		return status.Error(codes.Internal, errMessage)
	}
	return status.Error(code, errMessage)
}
