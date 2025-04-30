-- +goose Up
CREATE TABLE messages
(
    id         serial primary key,
    chat_id    int,
    creator    text,
    text       text,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
DROP TABLE messages;
