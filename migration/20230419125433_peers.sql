-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS peers
(
    nickname VARCHAR(25) PRIMARY KEY,
    birthday DATE NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE peers;
-- +goose StatementEnd
