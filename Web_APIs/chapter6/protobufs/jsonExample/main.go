package main

import (
	"fmt"

	"encoding/json"
	pb "github.com/web_apis/chapter6/protobufs/protofiles"
)

func main() {
	p := &pb.Person{
		Id: 1234,
		Email: "rf@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	body, _ := json.Marshal(p)
	fmt.Println(string(body))
}
