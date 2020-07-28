package managing

import (
	"context"
	"github.com/ktsymbal/seaports-catalogue/domain-service/proto"
	"log"
	"time"

	"github.com/ktsymbal/seaports-catalogue/pkg/entities"
	"google.golang.org/grpc"
)

const ConnectionTimeout = 5 * time.Second

type Service interface {
	GetPort(id string) (entities.Port, error)
	ListPorts(limit, offset uint) ([]entities.Port, error)
	CreatePort(entities.Port) error
	UpdatePort(port entities.Port) error
	CreateOrUpdatePorts(...entities.Port) error
	DeletePort(id string) error
}

type service struct {
	connection *grpc.ClientConn
	client     proto.RepositoryClient

	repoAddr string
	err      error
}

func NewService(repoAddress string) (Service, error) {
	log.Print("Managing service: dialing gRPC server...")
	ctx, cancel := context.WithTimeout(context.Background(), ConnectionTimeout)
	conn, err := grpc.DialContext(ctx, repoAddress, grpc.WithInsecure(), grpc.WithBlock())
	cancel()
	if err != nil {
		return nil, err
	}

	client := proto.NewRepositoryClient(conn)
	log.Print("Managing service: gRPC connection established.")
	return &service{repoAddr: repoAddress, connection: conn, client: client}, nil
}

func (s *service) Stop() error {
	if s.err != nil {
		return s.err
	}
	if s.connection != nil {
		log.Print("Managing service: Closing gRPC client connection...")
		if err := s.connection.Close(); err != nil {
			return err
		}
	}
	log.Print("Managing service stopped.")
	return nil
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
		// return just grpc code ? or just text ?
		return []entities.Port{}, err
	}
	ps := make([]entities.Port, 0, len(ports.Ports))
	for _, port := range ports.Ports {
		ps = append(ps, proto.ProtoToDomainPort(port))
	}
	return ps, nil
}

func (s *service) CreatePort(port entities.Port) error {
	p := proto.DomainToProtoPort(port)
	_, err := s.client.CreatePort(context.Background(), p)
	return err
}

func (s *service) UpdatePort(port entities.Port) error {
	_, err := s.client.UpdatePort(context.Background(), proto.DomainToProtoPort(port))
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

func (s *service) DeletePort(id string) error {
	_, err := s.client.Delete(context.Background(), &proto.Port{Id: id})
	return err
}
