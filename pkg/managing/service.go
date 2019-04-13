package managing

import (
	"context"

	"github.com/kateGlebova/seaports-catalogue/pkg/http/proto"

	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
	"google.golang.org/grpc"
)

type Service interface {
	GetPort(id string) (entities.Port, error)
	ListPorts(limit, offset uint) ([]entities.Port, error)
	CreatePort(entities.Port) error
	UpdatePort(id string) error
	CreateOrUpdatePorts(...entities.Port) error
}

func NewService(repoAddress string) Service {
	return &service{repoAddr: repoAddress}
}

type service struct {
	connection *grpc.ClientConn
	client     proto.RepositoryClient

	repoAddr string
	err      error
}

func (s service) GetPort(id string) (entities.Port, error) {
	port, err := s.client.GetPort(context.Background(), &proto.Port{Id: id})
	if err != nil {
		return entities.Port{}, err
	}
	return proto.ProtoToDomainPort(port), nil
}

func (s service) ListPorts(limit, offset uint) ([]entities.Port, error) {
	ports, err := s.client.ListPorts(context.Background(), &proto.ListRequest{Limit: uint64(limit), Offset: uint64(offset)})
	if err != nil {
		return []entities.Port{}, err
	}
	ps := make([]entities.Port, 0, len(ports.Ports))
	for _, port := range ports.Ports {
		ps = append(ps, proto.ProtoToDomainPort(port))
	}
	return ps, nil
}

func (s *service) CreatePort(port entities.Port) error {
	_, err := s.client.CreatePort(context.Background(), proto.DomainToProtoPort(port))
	return err
}

func (s *service) UpdatePort(id string) error {
	_, err := s.client.UpdatePort(context.Background(), &proto.Port{Id: id})
	return err
}

func (s *service) CreateOrUpdatePorts(ports ...entities.Port) error {
	ps := make([]*proto.Port, 0, len(ports))
	for _, port := range ports {
		ps = append(ps, proto.DomainToProtoPort(port))
	}
	_, err := s.client.CreateOrUpdatePorts(context.Background(), &proto.Ports{Ports: ps})
	return err
}
