package main

import (
	"github.com/psankar/grpc-compression-example/libs/go"
)

type Handler struct {
}

// func (h *Handler)

func main() {
	port := 8080

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	var h Handler

	go.RegisterExampleServiceServer(grpcServer, &h)
	log.Debugf("starting GRPC server at port: %d", port)
	grpcServer.Serve(lis)
}
