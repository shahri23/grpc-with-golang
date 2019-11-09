package main

import (
	"fmt"
	"log"
	"context"
	"github.com/shahri23/grpc-go-course/calculator/calpb"
	"google.golang.org/grpc"
)


func main () {
	fmt.Println("Hello I am a Calc Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err )
	}
	defer cc.Close()
	
	c := calpb.NewCalServiceClient(cc)
	fmt.Println("Created Client %v" , c )
doUnary(c)
}

func doUnary (c calpb.CalServiceClient) {
	fmt.Println("Starting a calculator rpc")
    req := &calpb.SumRequest{
		Number1: 13,
		Number2: 123,  	
	}

    res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("error calling sum rpc: %v", err )
	}
	log.Printf("response from sum is %v", res.Result)
}
 