package servers

import (
	"net"

	pb "github.com/dindasigma/go-microservices-user/packages/api/proto/users"

	"google.golang.org/grpc"
)

// GrpcServer implements a gRPC Server for the Order service
type GrpcServer struct {
	server   *grpc.Server
	errCh    chan error
	listener net.Listener
}

// NewGrpcServer is a convenience func to create a GrpcServer
func NewGrpcServer(service pb.UserServiceServer, port string) (GrpcServer, error) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return GrpcServer{}, err
	}
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, service)

	return GrpcServer{
		server:   server,
		listener: lis,
		errCh:    make(chan error),
	}, nil
}

// Start starts the server in the background, pushing any error to the error channel
func (g GrpcServer) Start() {
	go func() {
		g.errCh <- g.server.Serve(g.listener)
	}()
}

// Stop stops the gRPC server
func (g GrpcServer) Stop() {
	g.server.GracefulStop()
}

// Error returns the server's error channel
func (g GrpcServer) Error() chan error {
	return g.errCh
}
