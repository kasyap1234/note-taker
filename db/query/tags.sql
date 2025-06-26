-- name: CreateTag :one 
INSERT INTO tags(tag) VALUES ($1) RETURNING *; 
