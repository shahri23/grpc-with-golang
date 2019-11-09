package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"github.com/shahri23/grpc-go-course/calculator/calpb"
	"google.golang.org/grpc"
)

type server struct {}

func (*server) 	Sum(ctx context.Context, req *calpb.SumRequest) (*calpb.SumResponse, error) {
	fmt.Println("Recieved Sum request RPC: %v" , req)
	firstNumber := req.Number1;
	secondNumber := req.Number2;
	sum := firstNumber+secondNumber;
	res := &calpb.SumResponse{
		Result: sum,
	}
	return res, nil
}

func main () {
	fmt.Println("Calculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err )
	}

	s := grpc.NewServer()

	calpb.RegisterCalServiceServer(s, &server{})

	
	if err:=s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err )
	}

	}