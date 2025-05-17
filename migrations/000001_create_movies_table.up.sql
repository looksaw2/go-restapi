CREATE TABLE IF NOT EXISTS  movies (
    id bigserial PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);