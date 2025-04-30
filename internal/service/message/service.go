package message

import (
	"github.com/armanbektassov/go_chat/internal/client/db"
	"github.com/armanbektassov/go_chat/internal/repository"
	"github.com/armanbektassov/go_chat/internal/service"
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
