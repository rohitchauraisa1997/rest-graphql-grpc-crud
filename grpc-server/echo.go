package main

import (
	"context"
	"fmt"

	pb "github.com/rohitchauraisa1997/grpc-server/proto"
)

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	fmt.Println("-----------------------")
	fmt.Println("-----------------------")
	fmt.Println("-----------------------")
	fmt.Println("-----------------------")
	return req, nil
}
