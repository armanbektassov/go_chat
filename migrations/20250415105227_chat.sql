-- +goose Up
CREATE TABLE chats
(
    id         serial primary key,
    usernames  text,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose Down
DROP TABLE chat;
