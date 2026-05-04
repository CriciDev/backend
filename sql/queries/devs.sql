-- name: CreateDev :one
INSERT INTO devs (name, email, password, skills, bio, availability, socials) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING name, email, skills, bio, availability, socials, created_at;

-- name: FindDevByEmail :one
SELECT name, email, skills, bio, availability, socials, created_at FROM devs WHERE email = $1;