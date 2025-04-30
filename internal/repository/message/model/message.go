package model

import (
	"database/sql"
	"time"
)

type Message struct {
	ID        int64
	Info      MessageInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type MessageInfo struct {
	ChatId  int64
	Creator string
	Text    string
}
