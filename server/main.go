package main

import (
	"fmt"
	"log"
	"net"

	golibs "github.com/psankar/grpc-compression-example/libs/go"
	"google.golang.org/grpc"
)

type Handler struct {
	golibs.UnimplementedExampleServiceServer
}

func (h *Handler) Query(req *golibs.Request, stream golibs.ExampleService_QueryServer) error {
	return nil
}

func main() {
	port := 8080

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	var h Handler

	golibs.RegisterExampleServiceServer(grpcServer, &h)
	log.Printf("starting GRPC server at port: %d", port)
	grpcServer.Serve(lis)
}
