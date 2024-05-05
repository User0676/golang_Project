--create extension if not exists citext



--CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users
(
    id BIGSERIAL PRIMARY KEY,
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    name text NOT NULL,
    email CITEXT UNIQUE NOT NULL,
    password_hash bytea not  null,
    activated bool NOT NULL ,
    version INTEGER NOT NULL DEFAULT 1
);