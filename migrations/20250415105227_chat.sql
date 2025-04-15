-- +goose Up
CREATE TABLE chat (
                       id int NOT NULL,
                       creator text,
                       text text,
                       PRIMARY KEY(id)
);
-- +goose Down
DROP TABLE chat;
