package service

import (
	"context"
	"sync"

	// entity "github.com/muhammadchandra19/akira-go-chat/modules/chat/entity"
	"github.com/muhammadchandra19/akira-go-chat/modules/chat/entity"
	chatPb "github.com/muhammadchandra19/akira-go-proto/modules/chat/v1/public"
)

type ChatService struct {
	chatPb.UnimplementedChatServiceServer
	subscribers sync.Map
}

func (c *ChatService) CreateStream(connect *chatPb.StreamConnect, stream chatPb.ChatService_CreateStreamServer) error {
	// fin := make(chan bool)

	conn := &entity.Connection{
		Stream:   stream,
		Id:       connect.ChannelID,
		Finished: make(chan bool),
		Error:    make(chan error),
	}

	c.subscribers.Store(connect.GetChannelID(), conn)

	// return nil
	return <-conn.Error
}

func (c *ChatService) SendMessage(ctx context.Context, content *chatPb.ContentMessage) (*chatPb.Empty, error) {
	return &chatPb.Empty{}, nil
}

func (c *ChatService) CreateRoom(ctx context.Context, room *chatPb.Room) (*chatPb.Empty, error) {
	return &chatPb.Empty{}, nil
}

func (c *ChatService) AddUserToRoom(ctx context.Context, user *chatPb.UserRoom) (*chatPb.Empty, error) {
	return &chatPb.Empty{}, nil
}

func NewChatService(subscribers sync.Map) *ChatService {
	return &ChatService{subscribers: subscribers}
}
