-- name: CreateDev :one
INSERT INTO devs (name, email, password, skills, bio, availability, socials) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING name, email, skills, bio, availability, socials, created_at;

-- name: FindDevByEmail :one
SELECT id, name, email, password FROM devs WHERE email = $1;

-- name: FindDevByID :one
SELECT name, email, skills, bio, availability, socials FROM devs WHERE id = $1;

-- name: EmailAlreadyRegistered :one
SELECT count(*) AS total FROM devs WHERE email = $1;