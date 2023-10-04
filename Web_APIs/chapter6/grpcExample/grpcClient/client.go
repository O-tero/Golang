package main

import (
	"log"

	pb "github.com/web_apis/chapter6/grpcExample/Protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		// Create a client
		c := pb.NewMoneyTransactionClient(conn)
		from := "1234"
		to := "5678"
		amount := float32(1250.75)

		// Make a server request
		r, err := c.MakeTransaction(context.Background(),
		&pb.TransactionRequest{From: from,
			To: to, Amount: amount})
		if err != nil {
			log.Fatalf("Could not transact: %v", err)
		}
		log.Printf("Transaction confirmed: %t", r.Confirmation)
	}
}