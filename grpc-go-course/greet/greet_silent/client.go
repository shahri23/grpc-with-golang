package main

import (
	"fmt"
	"log"
	"context"
	"github.com/shahri23/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)


func main () {
	fmt.Println("Hello I am a Cleint")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err )
	}
	defer cc.Close()
	
	c := greetpb.NewGreetServiceClient(cc)
//	fmt.Println("Created Client %v" , c )
doUnary(c)
}

func doUnary (c greetpb.GreetServiceClient){
	fmt.Println("Starting a unary rpc")
    req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
		FirstName: "Shahzad",
		LastName: "Riasat",
		},
	}

    res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error calling greet rpc: %v", err )
	}
	log.Printf("response from greet is %v", res.Result)
}
 