CREATE TABLE devs (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    skills TEXT,
    bio TEXT,
    availability BOOLEAN NOT NULL DEFAULT true,
    socials TEXT[],
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);