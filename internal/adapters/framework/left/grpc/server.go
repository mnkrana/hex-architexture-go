package grpc

import (
	"log"
	"net"

	"github.com/mnkrana/hex/internal/adapters/framework/left/grpc/pb"
	"github.com/mnkrana/hex/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdaptor(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpcAdapter Adapter) Run() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen to port 9000: %v", err)
	}

	arithSvc := grpcAdapter
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithSvc)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}

}
