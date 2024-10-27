-- +goose Up

CREATE TABLE IF NOT EXISTS push_notifications (
    id SERIAL PRIMARY KEY,
    user_id UUID,
    notification TEXT,
    readed BOOLEAN,
    created_time TIMESTAMP,
    readed_time TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS push_notifications;