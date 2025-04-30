package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/armanbektassov/go_chat/internal/model"
	desc "github.com/armanbektassov/go_chat/pkg/message_v1"
)

func ToMessageFromService(message *model.Message) *desc.Message {
	var updatedAt *timestamppb.Timestamp
	if message.UpdatedAt.Valid {
		updatedAt = timestamppb.New(message.UpdatedAt.Time)
	}

	return &desc.Message{
		Id:        message.ID,
		Info:      ToMessageInfoFromService(message.Info),
		CreatedAt: timestamppb.New(message.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToMessageInfoFromService(info model.MessageInfo) *desc.MessageInfo {
	return &desc.MessageInfo{
		ChatId:  info.ChatId,
		Creator: info.Creator,
		Text:    info.Text,
	}
}

func ToMessageInfoFromDesc(info *desc.MessageInfo) *model.MessageInfo {
	return &model.MessageInfo{
		ChatId:  info.ChatId,
		Creator: info.Creator,
		Text:    info.Text,
	}
}
