package proto

import (
	"context"

	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
	"github.com/kateGlebova/seaports-catalogue/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RepositoryGRPCService struct {
	portRepo entities.PortRepository
}

func NewRepositoryGRPCService(portRepo entities.PortRepository) *RepositoryGRPCService {
	return &RepositoryGRPCService{portRepo: portRepo}
}

func (s RepositoryGRPCService) ListPorts(_ context.Context, r *ListRequest) (*Ports, error) {
	limit, offset := uint(r.Limit), uint(r.Offset)
	ports, err := s.portRepo.GetAllPorts(limit, offset)
	if err != nil {
		return &Ports{}, gRPCError(err)
	}

	ps := make([]*Port, 0, len(ports))
	for _, port := range ports {
		ps = append(ps, DomainToProtoPort(port))
	}
	return &Ports{Ports: ps}, nil
}

func (s RepositoryGRPCService) GetPort(_ context.Context, port *Port) (*Port, error) {
	p, err := s.portRepo.GetPort(port.Id)
	if err != nil {
		return port, gRPCError(err)
	}
	return DomainToProtoPort(p), err
}

func (s RepositoryGRPCService) CreatePort(_ context.Context, port *Port) (*Empty, error) {
	p := ProtoToDomainPort(port)
	err := s.portRepo.CreatePort(p)
	if err != nil {
		return &Empty{}, gRPCError(err)
	}
	return &Empty{}, nil
}

func (s RepositoryGRPCService) UpdatePort(_ context.Context, port *Port) (*Empty, error) {
	p := ProtoToDomainPort(port)
	err := s.portRepo.UpdatePort(p)
	if err != nil {
		return &Empty{}, gRPCError(err)
	}
	return &Empty{}, nil
}

func (s RepositoryGRPCService) CreateOrUpdatePorts(_ context.Context, ports *Ports) (*Empty, error) {
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
	code, ok := storage.GRPCErrorMapping[errMessage]
	if !ok {
		return status.Error(codes.Internal, errMessage)
	}
	return status.Error(code, errMessage)
}
