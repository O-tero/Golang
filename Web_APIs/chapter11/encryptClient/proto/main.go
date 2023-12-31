package main

import (
	"context"
	"fmt"

	proto "github.com/git-user/chapter11/encryptClient/proto"
	micro "github.com/micro/go-micro"
	"github.com/web_apis/chapter11/encryptClient/proto"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("encrypter.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new encrypter service instance
	encrypter := proto.NewEncrypterService("encrypter", service.Client())

	// Call the encrypter
	rsp, err := encrypter.Encrypt(context.TODO(), &proto.Request{
		Message : "I am a Message",
		Key: "111023043350789514532147",
	})

	if err != nil {
		fmt.Println(err)
	}

	// Print response 
	fmt.Println(rsp.Result)

	// Call the decrypter
	rsp, err := encrypter.Decrypt(context.TODO(), &proto.Request {
		Message: rsp.Result,
		Key: "111023043350789514532147",
	}) 
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Result)
}
