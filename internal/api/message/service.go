package message

import (
	"github.com/armanbektassov/go_chat/internal/service"
	desc "github.com/armanbektassov/go_chat/pkg/message_v1"
)

type Implementation struct {
	desc.UnimplementedMessageV1Server
	messageService service.MessageService
}

func NewImplementation(messageService service.MessageService) *Implementation {
	return &Implementation{
		messageService: messageService,
	}
}
