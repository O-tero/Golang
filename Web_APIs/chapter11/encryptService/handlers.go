package main

import (
	"context"

	proto "github.com/web_apis/chapter11/encryptService/proto"
)

// Encrypter holds the information about methods
type Encrypter struct{}

// Encrypt converts a message into cipher and return response

func (g *Encrypter) Encrypt(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Result = EncryptString(req.Key, req.Message)
	return nil
}

// Decrypt converts a cipher into message and return response
func (g *Encrypter) Decrypt(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Result = DecryptString(req.Key, req.Message)
	return nil
}
