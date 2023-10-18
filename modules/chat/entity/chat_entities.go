package entity

import (
	chatPb "github.com/muhammadchandra19/akira-go-proto/modules/chat/v1/public"
)

type Connection struct {
	Stream   chatPb.ChatService_CreateStreamServer
	Id       string
	Finished chan bool
	Error    chan error
}
