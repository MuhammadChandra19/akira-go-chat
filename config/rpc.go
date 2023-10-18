package config

import (
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"

	chatConfig "github.com/muhammadchandra19/akira-go-chat/modules/chat/config"
)

type GRPCServer struct {
	Server *grpc.Server
}

func (gr GRPCServer) Stop() {
	gr.Server.GracefulStop()
}

func registerRPCServer(server *grpc.Server, rdsClient *redis.Client) {
	chatConfig.RegisterChatRPCServer(server, rdsClient)
}

func NewGRPCServer() (*GRPCServer, error) {
	grpcServer := grpc.NewServer()
	rdsClient := RedisClient()
	registerRPCServer(grpcServer, rdsClient)

	return &GRPCServer{
		Server: grpcServer,
	}, nil
}
