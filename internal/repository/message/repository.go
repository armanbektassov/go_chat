package message

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/armanbektassov/go_chat/internal/model"
	"github.com/armanbektassov/go_chat/internal/repository"
	"github.com/armanbektassov/go_chat/internal/repository/message/converter"
	modelRepo "github.com/armanbektassov/go_chat/internal/repository/message/model"
	"github.com/armanbektassov/platform_common/pkg/client/db"
)

const (
	tableName = "messages"

	idColumn        = "id"
	chatIdColumn    = "chat_id"
	creatorColumn   = "creator"
	textColumn      = "text"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.MessageRepository {
	return &repo{db: db}
}

func (r *repo) SendMessage(ctx context.Context, info *model.MessageInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIdColumn, creatorColumn, textColumn).
		Values(info.ChatId, info.Creator, info.Text).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "message_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Message, error) {
	builder := sq.Select(idColumn, chatIdColumn, creatorColumn, textColumn, createdAtColumn, updatedAtColumn).
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

	var message modelRepo.Message
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&message.ID, &message.Info.ChatId, &message.Info.Creator, &message.Info.Text, &message.CreatedAt, &message.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToMessageFromRepo(&message), nil
}
