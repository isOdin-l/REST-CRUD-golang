CREATE TABLE users
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(64) not null,
    username varchar(64) not null unique,
    password_hash TEXT not null
);

CREATE TABLE todo_lists
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title varchar(128) not null,
    description TEXT
);

CREATE TABLE users_lists
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID references users(id) on delete cascade not null,
    list_id UUID references todo_lists(id) on delete cascade not null
);

CREATE TABLE todo_items
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- delete uuid genereation
    title varchar(128) not null,
    description TEXT,
    done boolean not null default false
);

CREATE TABLE lists_items
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    item_id UUID references todo_items(id) on delete cascade not null,
    list_id UUID references todo_lists(id) on delete cascade not null
);