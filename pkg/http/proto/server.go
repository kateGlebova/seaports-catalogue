package proto

import (
	"log"
	"net"

	"github.com/kateGlebova/seaports-catalogue/pkg/entities"

	"github.com/kateGlebova/seaports-catalogue/pkg/lifecycle"
	"google.golang.org/grpc"
)

type PortDomainService struct {
	grpcService RepositoryServer
	server      *grpc.Server

	port string
	err  error
}

func NewPortDomainService(port string, storage entities.PortRepository) *PortDomainService {
	grpcServer := grpc.NewServer()
	grpcService := NewRepositoryGRPCService(storage)
	RegisterRepositoryServer(grpcServer, grpcService)
	return &PortDomainService{grpcService: grpcService, port: port, server: grpcServer}
}

// Run starts PortDomainService gRPC server
func (r *PortDomainService) Run() {
	log.Printf("Listening on %s...", r.port)
	lis, err := net.Listen("tcp", ":"+r.port)
	if err != nil {
		r.err = err
		lifecycle.KillTheApp()
	}

	if err = r.server.Serve(lis); err != grpc.ErrServerStopped {
		r.err = err
		lifecycle.KillTheApp()
	}
}

// Stop gracefully stops gRPC server
func (r *PortDomainService) Stop() (err error) {
	log.Print("PortDomainService is gracefully stopping...")
	if r.err != nil {
		return r.err
	}
	if r.server != nil {
		r.server.GracefulStop()
		log.Print("PortDomainService stopped")
	}
	return
}
