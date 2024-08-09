-- +goose Up

CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    user_id UUID,
    notification TEXT,
    readed BOOLEAN,
    created_time TIMESTAMP,
    readed_time TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS notifications;