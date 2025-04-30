package chat

import (
	"context"
	"log"

	"github.com/armanbektassov/go_chat/internal/converter"
	desc "github.com/armanbektassov/go_chat/pkg/chat_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.chatService.Create(ctx, converter.ToChatFromDesc(req.GetUsernames()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted chat with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
