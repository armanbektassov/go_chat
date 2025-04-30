package app

import (
	"context"
	"log"

	"github.com/armanbektassov/go_chat/internal/api/chat"
	"github.com/armanbektassov/go_chat/internal/api/message"
	"github.com/armanbektassov/go_chat/internal/config"
	"github.com/armanbektassov/go_chat/internal/repository"
	chatRepository "github.com/armanbektassov/go_chat/internal/repository/chat"
	messageRepository "github.com/armanbektassov/go_chat/internal/repository/message"
	"github.com/armanbektassov/go_chat/internal/service"
	chatService "github.com/armanbektassov/go_chat/internal/service/chat"
	messageService "github.com/armanbektassov/go_chat/internal/service/message"
	"github.com/armanbektassov/platform_common/pkg/client/db"
	"github.com/armanbektassov/platform_common/pkg/client/db/pg"
	"github.com/armanbektassov/platform_common/pkg/client/db/transaction"
	"github.com/armanbektassov/platform_common/pkg/closer"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient          db.Client
	txManager         db.TxManager
	chatRepository    repository.ChatRepository
	messageRepository repository.MessageRepository

	chatService    service.ChatService
	messageService service.MessageService

	chatImpl    *chat.Implementation
	messageImpl *message.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = messageRepository.NewRepository(s.DBClient(ctx))
	}

	return s.messageRepository
}

func (s *serviceProvider) MessageService(ctx context.Context) service.MessageService {
	if s.messageService == nil {
		s.messageService = messageService.NewService(
			s.MessageRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.messageService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}

func (s *serviceProvider) MessageImpl(ctx context.Context) *message.Implementation {
	if s.messageImpl == nil {
		s.messageImpl = message.NewImplementation(s.MessageService(ctx))
	}

	return s.messageImpl
}
