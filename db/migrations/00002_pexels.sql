-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS note_images(
id SERIAL PRIMARY KEY, 
note_id INT REFERENCES notes(id) ON DELETE CASCADE, 
image_url TEXT NOT NULL, 
caption TEXT,
created_at TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE note_images;
-- +goose StatementEnd
