-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS notes(
id SERIAL PRIMARY KEY, 
title TEXT NOT NULL, 
content TEXT NOT NULL, 
created_at TIMESTAMPTZ DEFAULT NOW()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE notes;
-- +goose StatementEnd
