package chat

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	modelRepo "github.com/armanbektassov/go_chat/internal/repository/chat/model"

	"github.com/armanbektassov/go_chat/internal/model"
	"github.com/armanbektassov/go_chat/internal/repository"
	"github.com/armanbektassov/go_chat/internal/repository/chat/converter"
	"github.com/armanbektassov/platform_common/pkg/client/db"
)

const (
	tableName = "chats"

	idColumn        = "id"
	usernamesColumn = "usernames"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(usernamesColumn).
		Values(chat.Usernames).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Chat, error) {
	builder := sq.Select(idColumn, usernamesColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var chat modelRepo.Chat
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chat.ID, &chat.Usernames, &chat.CreatedAt, &chat.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToChatFromRepo(&chat), nil
}
