package converter

import (
	"github.com/armanbektassov/go_chat/internal/model"
	modelRepo "github.com/armanbektassov/go_chat/internal/repository/chat/model"
)

func ToChatFromRepo(chat *modelRepo.Chat) *model.Chat {
	return &model.Chat{
		ID:        chat.ID,
		Usernames: chat.Usernames,
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}
}
