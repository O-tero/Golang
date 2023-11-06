package main

import (
	fmt "fmt"

	proto "github.com/web_apis/chapter11/encryptService/proto"
	micro "github.com/micro/go-micro"
)

func main() {
	// Create a new service.Optionally include some options here
	service := micro.NewService(
		micro.Name("encrypter"),
	)

	// Init wil parse the command line flags
	service.Init()

	// Register handler
	proto.RegisterEncrypterHandler(service.Server(), new(Encrypter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
