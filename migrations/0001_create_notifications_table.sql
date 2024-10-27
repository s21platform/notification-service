-- +goose Up

CREATE TABLE IF NOT EXISTS push_notifications (
    id SERIAL PRIMARY KEY,
    user_id UUID,
    notification TEXT,
    is_read BOOLEAN,
    created_time TIMESTAMP,
    read_time TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS push_notifications;