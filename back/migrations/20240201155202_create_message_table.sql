-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS message (
    id INT AUTO_INCREMENT,
    content TEXT NOT NULL,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    conversation_id INT NOT NULL,
    is_user BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (id),
    FOREIGN KEY (conversation_id) REFERENCES conversation(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS message;
-- +goose StatementEnd