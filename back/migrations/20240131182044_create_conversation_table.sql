-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS conversation (
    id INT AUTO_INCREMENT,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS conversation;
-- +goose StatementEnd