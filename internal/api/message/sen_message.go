package message

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"

	"github.com/armanbektassov/go_chat/internal/converter"
	desc "github.com/armanbektassov/go_chat/pkg/message_v1"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	_, err := i.messageService.SendMessage(ctx, converter.ToMessageInfoFromDesc(req.GetMessageInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted message")

	return &emptypb.Empty{}, nil
}
