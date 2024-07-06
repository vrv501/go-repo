package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/vrv501/go-repo/grpc/generated/bank"
	app "github.com/vrv501/go-repo/grpc/internal/app"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	bs := app.NewBankService()
	reflection.Register(server)
	pb.RegisterBankServiceServer(server, bs)

	log.Println("listening on 50051")
	if err = server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
