package managing

import (
	"log"

	"github.com/kateGlebova/seaports-catalogue/pkg/http/proto"
	"github.com/kateGlebova/seaports-catalogue/pkg/lifecycle"
	"google.golang.org/grpc"
)

func (s *service) Run() {
	conn, err := grpc.Dial(s.repoAddr, grpc.WithInsecure())
	if err != nil {
		s.err = err
		lifecycle.KillTheApp()
	}
	s.connection = conn

	client := proto.NewRepositoryClient(conn)
	s.client = client
}

func (s *service) Stop() error {
	if s.err != nil {
		return s.err
	}
	log.Print("Closing gRPC client connection...")
	if s.connection != nil {
		return s.connection.Close()
	}
	return nil
}
