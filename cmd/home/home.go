package main

import (
	"context"
	"log"
	"net"

	hspb "github.com/chokevin/home/homeserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	hspb.UnimplementedHomeServiceServer
}

var properties = []*hspb.GetPropertyResponse{
	{Id: 1, Address: "123 Aloha St", City: "Honolulu", State: "HI", ZipCode: "96815", Price: 950000, Bedrooms: 3, Bathrooms: 2, SquareFeet: 1200},
	{Id: 2, Address: "456 Beach Blvd", City: "Maui", State: "HI", ZipCode: "96753", Price: 1200000, Bedrooms: 4, Bathrooms: 3, SquareFeet: 1800},
}

func (s *server) GetProperty(ctx context.Context, req *hspb.GetPropertyRequest) (*hspb.GetPropertyResponse, error) {
	for _, property := range properties {
		if property.Id == req.Id {
			return property, nil
		}
	}
	return nil, grpc.Errorf(grpc.Code(hspb.GetPropertyResponse{}), "Property not found")
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hspb.RegisterHomeServiceServer(s, &server{})
	reflection.Register(s)

	log.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
