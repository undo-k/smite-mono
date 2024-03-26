package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/undo-k/smite-mono/protos/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AggregatorServer struct {
	protos.UnimplementedAggregatorServer
	app *Config
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))

	if err != nil {
		log.Fatalf("Failed to listen to grpc: %v", err)
	}

	s := grpc.NewServer()

	protos.RegisterAggregatorServer(s, &AggregatorServer{app: app})

	log.Printf("gRPC Server started on port %s", grpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed serve grpc: %v", err)
	}

}

func (a *AggregatorServer) FetchData(ctx context.Context, req *protos.AggregateRequest) (*protos.AggregateResponse, error) {

	go batchData(int(req.NumberOfRequests))

	return &protos.AggregateResponse{
		Ok: true,
	}, nil

}

func PutGodViaGRPC(god *protos.God) error {
	conn, err := grpc.Dial("cache-service:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	defer conn.Close()

	client := protos.NewGodCacheClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = client.PutGod(ctx, god)

	if err != nil {
		return err
	}

	return nil
}
