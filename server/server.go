package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	pb "currency-converter/currency"
)

type server struct{
	pb.UnimplementedCurrencyConverterServer
}

func (s *server) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	rates := map[string]float32{
		"USD": 1.0,
		"EUR": 0.85,
		"JPY": 110.0,
	}

	fromCurrency := req.GetFromCurrency()
	toCurrency := req.GetToCurrency()
	amount := req.GetAmount()

	log.Printf("Received from: %v, to: %v, amount: %v", fromCurrency, toCurrency, amount)

	fromRate, ok := rates[fromCurrency]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "Unknown from currency: %s", fromCurrency)
	}

	toRate, ok := rates[toCurrency]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "Unknown to currency: %s", toCurrency)
	}

	convertedAmount := amount * fromRate / toRate
	
	return &pb.ConvertResponse{ConvertedAmount: convertedAmount}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCurrencyConverterServer(s, &server{})
	log.Println("gRPC server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}