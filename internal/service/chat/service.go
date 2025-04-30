package chat

import (
	"github.com/armanbektassov/go_chat/internal/repository"
	"github.com/armanbektassov/go_chat/internal/service"
	"github.com/armanbektassov/platform_common/pkg/client/db"
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
