package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/armanbektassov/go_chat/internal/model"
	desc "github.com/armanbektassov/go_chat/pkg/chat_v1"
)

func ToChatFromService(chat *model.Chat) *desc.Chat {
	var updatedAt *timestamppb.Timestamp
	if chat.UpdatedAt.Valid {
		updatedAt = timestamppb.New(chat.UpdatedAt.Time)
	}

	return &desc.Chat{
		Id:        chat.ID,
		Usernames: chat.Usernames,
		CreatedAt: timestamppb.New(chat.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToChatFromDesc(usernames string) *model.Chat {
	return &model.Chat{
		Usernames: usernames,
	}
}
