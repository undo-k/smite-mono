package main

import (
	"context"
	"fmt"
	"time"

	"github.com/undo-k/smite-mono/protos/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (app *Config) FetchGodViaGRPC(godName string) (*protos.God, error) {

	conn, err := grpc.Dial("cache-service:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := protos.NewGodCacheClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	gr, err := client.FetchGod(ctx, &protos.GodRequest{
		Name: godName,
	},
	)

	if err != nil {
		return nil, err
	}

	return gr, nil
}

func (app *Config) FetchGodListViaGRPC() (*protos.GodList, error) {
	conn, err := grpc.Dial("cache-service:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := protos.NewGodCacheClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	godlist, err := client.FetchAllGods(ctx, &protos.GodRequest{
		Name: "all",
	},
	)

	if err != nil {
		return nil, err
	}

	return godlist, nil
}

func (app *Config) PutGodViaGRPC(god *protos.God) error {
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

func (app *Config) triggerAggregatorViaGRPC(numberOfRequests int) error {
	conn, err := grpc.Dial("aggregator:5002", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not dial aggregator:5002")
		return err
	}

	defer conn.Close()

	client := protos.NewAggregatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = client.FetchData(ctx, &protos.AggregateRequest{NumberOfRequests: int32(numberOfRequests)})

	if err != nil {
		fmt.Println("error calling FetchData")
		return err
	}

	return nil
}
