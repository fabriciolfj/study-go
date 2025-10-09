package main

import (
	"context"
	"log"
	"net"

	pb "hello/proto" // ajuste o import

	"google.golang.org/grpc"
)

type server struct {
	pb.ChatServiceServer
}

func (s *server) RouteComments(ctx context.Context, req *pb.CommentRequest) (*pb.CommentResponse, error) {
	return &pb.CommentResponse{
		CommentLength:        1000,
		PreviousCommentCount: 2000,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})

	log.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
