-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS note_tags(
note_id INT REFERENCES notes(id) ON DELETE CASCADE, 
tag_id INT REFERENCES tags(id) ON DELETE CASCADE, 
PRIMARY KEY (note_id,tag_id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE note_tags;
-- +goose StatementEnd
