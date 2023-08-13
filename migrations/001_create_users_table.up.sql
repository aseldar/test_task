-- 001_create_users_table.up.sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    firstname VARCHAR(100),
    lastname VARCHAR(100),
    email VARCHAR(100),
    age INT,
    created TIMESTAMPTZ
);
