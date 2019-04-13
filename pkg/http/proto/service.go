package proto

import (
	"context"

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

func (s RepositoryService) ListPorts(_ context.Context, r *ListRequest) (*Ports, error) {
	limit, offset := uint(r.Limit), uint(r.Offset)
	if limit == 0 {
		limit = DefaultLimit
	}
	if offset == 0 {
		offset = DefaultOffset
	}

	ports := make([]*Port, 0, limit)

	for _, port := range s.portRepo.GetAllPorts(limit, offset) {
		ports = append(ports, DomainToProtoPort(port))
	}
	return &Ports{Ports: ports}, nil
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

func (s RepositoryService) CreateOrUpdatePorts(_ context.Context, ports *Ports) (*Empty, error) {
	ps := make([]entities.Port, 0, len(ports.Ports))
	for _, port := range ports.Ports {
		ps = append(ps, ProtoToDomainPort(port))
	}

	err := s.portRepo.CreateOrUpdatePorts(ps...)
	if err != nil {
		return &Empty{}, gRPCError(err)
	}
	return &Empty{}, nil
}

func gRPCError(err error) error {
	errMessage := err.Error()
	code, ok := errorMapping[errMessage]
	if !ok {
		return status.Error(codes.Internal, errMessage)
	}
	return status.Error(code, errMessage)
}
