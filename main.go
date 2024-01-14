package main

import (
	"log"
	"net"

	"ticketing_app/api"
	ticket "ticketing_app/proto-gen/ticket"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":4080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	ticket_server := api.NewTrainService()

	ticket.RegisterTrainTicketServiceServer(s, ticket_server)
	reflection.Register(s)

	log.Println("Server listening on :4080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
