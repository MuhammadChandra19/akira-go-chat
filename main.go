package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/muhammadchandra19/akira-go-chat/config"
	"google.golang.org/grpc"
)

var addr string = "localhost:5051"

func main() {

	grpcServer, err := config.NewGRPCServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func(gs *grpc.Server, lis net.Listener) {
		log.Printf("grpc Server is available at %s\n", addr)
		if err = gs.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}(grpcServer.Server, lis)

	<-sigChan

	log.Println("Shutting down the grpc Server...")

	grpcServer.Stop()
	log.Println("grpc Server gracefully stopped")
}
