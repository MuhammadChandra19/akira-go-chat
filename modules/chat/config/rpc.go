package config

import (
	"sync"

	"github.com/go-redis/redis/v8"
	service "github.com/muhammadchandra19/akira-go-chat/modules/chat/services"
	chatPb "github.com/muhammadchandra19/akira-go-proto/modules/chat/v1/public"
	"google.golang.org/grpc"
)

func RegisterChatRPCServer(server *grpc.Server, rds *redis.Client) {
	mapConn := sync.Map{}
	chatService := service.NewChatService(mapConn)

	chatPb.RegisterChatServiceServer(server, chatService)
}
