package converter

import (
	"github.com/armanbektassov/go_chat/internal/model"
	modelRepo "github.com/armanbektassov/go_chat/internal/repository/message/model"
)

func ToMessageFromRepo(message *modelRepo.Message) *model.Message {
	return &model.Message{
		ID:        message.ID,
		Info:      ToMessageInfoFromRepo(message.Info),
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}
}
func ToMessageInfoFromRepo(info modelRepo.MessageInfo) model.MessageInfo {
	return model.MessageInfo{
		ChatId:  info.ChatId,
		Creator: info.Creator,
		Text:    info.Text,
	}
}
