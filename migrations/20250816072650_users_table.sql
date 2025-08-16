-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
     id VARCHAR(36) PRIMARY KEY,
     login text NOT NULL UNIQUE,
     password text NOT NULL,
     active boolean default true,
     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
     updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
