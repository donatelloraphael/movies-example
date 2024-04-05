-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  name text NOT NULL,
  email text UNIQUE NOT NULL,
  password_hash bytea NOT NULL,
  activated bool NOT NULL DEFAULT FALSE,
  version integer NOT NULL DEFAULT 1
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
