-- name: CreateImage :one 
INSERT INTO note_images(note_id,image_url,caption) VALUES($1,$2,$3) RETURNING *; 


-- name: GetImagesByNoteID :many 
SELECT * FROM note_images where note_id=$1 ORDER BY created_at ASC; 


-- name: GetImageByID :one 
SELECT * FROM note_images where id=$1; 
