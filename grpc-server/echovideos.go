package main

import (
	"context"
	"fmt"

	pb "github.com/rohitchauraisa1997/grpc-server/proto"
)

func (v *videoServer) Echo(ctx context.Context, req *pb.TestResponseRequest) (*pb.TestResponseRequest, error) {
	fmt.Println("*************************")
	return req, nil
}
