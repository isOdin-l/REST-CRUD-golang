CREATE TABLE users
(
    id UUID PRIMARY KEY,
    name varchar(64) NOT NULL,
    username varchar(64) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);

CREATE TABLE lists
(
    id UUID PRIMARY KEY,
    author_id UUID REFERENCES users(id) ON DELETE CASCADE,
    title varchar(128) NOT NULL,
    description TEXT
);

CREATE TABLE items
(
    id UUID PRIMARY KEY,
    list_id REFERENCES lists(id) ON DELETE CASCADE,
    title VARCHAR(128) NOT NULL,
    description TEXT,
    done BOOLEAN NOT NULL DEFAULT false
);