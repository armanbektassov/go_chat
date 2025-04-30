package chat

import (
	"github.com/armanbektassov/go_chat/internal/service"
	desc "github.com/armanbektassov/go_chat/pkg/chat_v1"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
