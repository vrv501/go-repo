package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/vrv501/go-repo/grpc/generated/bank"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	/*
		Max timeout per rpc 120s
		If we got status unavailable or resource exhausted we will retry(max number = 3)
		After first call, 2nd call will be made after random(0, initialBackoff)s
		Then after every nth(>=2) attempt will be made after random(0, min(initialBackoff * (backoffMultiplier)**(n - 1), maxBackoff))s

		But lets say during retry we got some another statusCode, we immediately stop
	*/
	Config = `
	{
		"methodConfig": [
			{
				"name": [{"service": "bank.v1.BankService"}],
				"waitForReady": true,
				"timeout": "120.0s",
				"retryPolicy": {
					"maxAttempts": 3,
					"initialBackoff": "5s",
					"maxBackoff": "10s",
					"backoffMultiplier": 1.0,
					"retryableStatusCodes": ["UNAVAILABLE", "RESOURCE_EXHAUSTED"]
				}
			}
		],
		"healthCheckConfig": {
			"serviceName": "bank.v1.BankService"
		}
	}
	`
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		grpc.WithDefaultServiceConfig(Config),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewBankServiceClient(conn)
	/*
		Since we have set timeout in serviceConfig, it's no longer required to pass via context explicitly
		If passed, min(svcConfigValue, ctxValue) will be considered
	*/

	// you might be wondering why cant we just add kv-pairs to ctx since ctx is being used
	// as parameter to reqSend func. ctx values are not passed over the wire.
	// in grpc-go its only used to propagate deadlines, signals
	md := metadata.New(map[string]string{"authorization": "dummy"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, err := client.FetchExchangeRates(ctx, &pb.ExchangeRateRequest{
		FromCurr: "asdasdasd",
		ToCurr:   "asdasdasd",
	}, grpc.Header(&md))
	if err != nil {
		log.Fatal("func", err)
	}

	if err != nil {
		log.Fatal("header ", err)
	}

	stream.Header()
	for {
		r, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			s, ok := status.FromError(err)
			if ok {
				fmt.Println(s.Code(), s.Message())
			}
			log.Fatal(err)
		}

		md, ok := metadata.FromIncomingContext(stream.Context())
		if ok {
			fmt.Println(md.Get("ASdasd"))
		}
		fmt.Println(r)
	}

}
