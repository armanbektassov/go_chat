package chat

import (
	"github.com/armanbektassov/go_chat/internal/client/db"
	"github.com/armanbektassov/go_chat/internal/repository"
	"github.com/armanbektassov/go_chat/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

func NewService(
	chatRepository repository.ChatRepository,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}
