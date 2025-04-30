package message

import (
	"github.com/armanbektassov/go_chat/internal/repository"
	"github.com/armanbektassov/go_chat/internal/service"
	"github.com/armanbektassov/platform_common/pkg/client/db"
)

type serv struct {
	messageRepository repository.MessageRepository
	txManager         db.TxManager
}

func NewService(
	messageRepository repository.MessageRepository,
	txManager db.TxManager,
) service.MessageService {
	return &serv{
		messageRepository: messageRepository,
		txManager:         txManager,
	}
}
