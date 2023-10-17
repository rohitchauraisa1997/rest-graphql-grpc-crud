package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/rohitchauraisa1997/grpc-server/proto"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.TestApiServer
}

type videoServer struct {
	pb.VideoApiServer
}

func main() {
	go func() {
		// mux
		mux := runtime.NewServeMux()

		// register the server
		pb.RegisterTestApiHandlerServer(context.TODO(), mux, &testApiServer{})
		pb.RegisterVideoApiHandlerServer(context.TODO(), mux, &videoServer{})

		// http server
		// log.Fatalln(http.ListenAndServe("localhost:8081", mux))
		// CORS Handler
		corsHandler := cors.New(cors.Options{
			AllowedOrigins: []string{"*"}, // You may restrict this to specific origins
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"*"}, // You may restrict this to specific headers
		})

		// Wrap the ServeMux with the CORS handler
		handler := corsHandler.Handler(mux)

		// http server
		// reason for 0.0.0.0 instead of localhost 
		// https://chat.openai.com/share/210446c0-6315-4723-b62e-d5d89b747f60
		log.Fatalln(http.ListenAndServe("0.0.0.0:8081", handler))

	}()

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, &testApiServer{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}
