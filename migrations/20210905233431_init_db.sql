-- +goose Up
-- +goose StatementBegin
CREATE TABLE animal(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name TEXT NOT NULL,
    type TEXT NOT NULL DEFAULT 'UNKNOWN'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE animal;
-- +goose StatementEnd
