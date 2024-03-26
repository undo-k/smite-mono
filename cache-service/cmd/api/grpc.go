package main

import (
	"context"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/undo-k/smite-mono/protos/protos"
	"google.golang.org/grpc"
)

type CacheServer struct {
	protos.UnimplementedGodCacheServer
	app *Config
}

func (c *CacheServer) FetchGod(ctx context.Context, req *protos.GodRequest) (*protos.God, error) {
	godName := req.GetName()

	cachedGod, err := c.app.retrieveFromCache(godName)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return cachedGod, nil
}

func (c *CacheServer) FetchAllGods(ctx context.Context, req *protos.GodRequest) (*protos.GodList, error) {
	var godList []*protos.God

	for _, god := range c.app.GodCache {
		godList = append(godList, god)
	}

	return &protos.GodList{Gods: godList}, nil
}

func (c *CacheServer) PutGod(ctx context.Context, god *protos.God) (*protos.Response, error) {

	println("Attempting to PUT a GOD")

	c.app.cacheInsert(god)

	return &protos.Response{Ok: true}, nil

}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen to grpc: %v", err)
	}

	s := grpc.NewServer()

	protos.RegisterGodCacheServer(s, &CacheServer{app: app})

	log.Printf("gRPC Server started on port %s", grpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed serve grpc: %v", err)
	}

}
