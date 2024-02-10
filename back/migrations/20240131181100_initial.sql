-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user (
    id INT AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user;
-- +goose StatementEnd