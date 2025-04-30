package repository

import (
	"context"

	"github.com/armanbektassov/go_chat/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context, info *model.Chat) (int64, error)
	Get(ctx context.Context, id int64) (*model.Chat, error)
}

type MessageRepository interface {
	SendMessage(ctx context.Context, info *model.MessageInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Message, error)
}
