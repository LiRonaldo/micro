package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	ServicesGW "micro/servicesGW"
	"net/http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	gRPCEndPonit := "localhost:8001"
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := ServicesGW.RegisterTestServiceHandlerFromEndpoint(ctx, mux, gRPCEndPonit, opts)
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe("localhost:9000", mux)
}
