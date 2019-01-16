package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	pb "../customer"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type CustomerServer struct {
	savedCustomers []*pb.CustomerRequest
}

func (s *CustomerServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("fff")
	fmt.Println(in)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
func (s *CustomerServer) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	return &pb.CustomerResponse{Id: 0}, nil
}

// GetCustomers returns all customers by given filter
func (s *CustomerServer) GetCustomers(filter *pb.CustomerFilter, stream pb.Customer_GetCustomersServer) error {
	for _, customer := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(customer); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCustomerServer(s, &CustomerServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
