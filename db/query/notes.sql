-- name: CreateNote :one 
INSERT INTO notes (title,content) VALUES ($1,$2) RETURNING *; 

-- name: GetNoteByID :one 
SELECT * FROM notes WHERE id=$1; 

-- name: ListNotes :many 
SELECT * FROM notes ORDER BY created_at desc; 
