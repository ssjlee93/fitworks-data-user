-- Migration: 001_create_users.sql
-- Creates the user_role enum and the users table.

CREATE TYPE user_role AS ENUM ('trainer', 'trainee', 'none');

CREATE TABLE IF NOT EXISTS users (
    id         BIGSERIAL    PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name  VARCHAR(100) NOT NULL,
    email      VARCHAR(255) NOT NULL UNIQUE,
    role       user_role    NOT NULL DEFAULT 'none',
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

-- Optional: index on email for fast lookup.
CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);
