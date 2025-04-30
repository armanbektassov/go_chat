package message

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/armanbektassov/go_chat/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, message *model.MessageInfo) (*emptypb.Empty, error) {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		_, errTx = s.messageRepository.SendMessage(ctx, message)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
