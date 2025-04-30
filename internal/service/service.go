package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/armanbektassov/go_chat/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
}

type MessageService interface {
	SendMessage(ctx context.Context, chat *model.MessageInfo) (*emptypb.Empty, error)
}
