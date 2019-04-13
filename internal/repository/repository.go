package repository

import (
	"log"
	"net"

	"github.com/kateGlebova/seaports-catalogue/pkg/http/proto"
	"github.com/kateGlebova/seaports-catalogue/pkg/shutdown"
	"google.golang.org/grpc"
)

type PortDomainService struct {
	grpcService proto.RepositoryServer
	server      *grpc.Server

	port string
	err  error
}

func NewPortDomainService(grpcService proto.RepositoryServer, port string) *PortDomainService {
	grpcServer := grpc.NewServer()
	proto.RegisterRepositoryServer(grpcServer, grpcService)
	return &PortDomainService{grpcService: grpcService, port: port, server: grpcServer}
}

// Run starts PortDomainService gRPC server
func (r *PortDomainService) Run() {
	log.Printf("Listening on %s...", r.port)
	lis, err := net.Listen("tcp", ":"+r.port)
	if err != nil {
		r.err = err
		shutdown.KillTheApp()
	}

	if err = r.server.Serve(lis); err != grpc.ErrServerStopped {
		r.err = err
		shutdown.KillTheApp()
	}
}

// Stop gracefully stops gRPC server
func (r *PortDomainService) Stop() (err error) {
	if r.err != nil {
		return r.err
	}
	if r.server != nil {
		r.server.GracefulStop()
		log.Print("PortDomainService stopped")
	}
	return
}
