-- +goose Up
-- +goose StatementBegin
CREATE TABLE profiles (
       id VARCHAR(36) PRIMARY KEY,
       user_id VARCHAR(36) NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
       first_name VARCHAR(100),
       last_name VARCHAR(100),
       phone VARCHAR(20),
       gender VARCHAR(10),
       created_at TIMESTAMP NOT NULL DEFAULT NOW(),
       updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_profiles_user_id ON profiles(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table profiles;
-- +goose StatementEnd
