CREATE TABLE IF NOT EXISTS users
(
    id UUID         NOT NULL PRIMARY KEY,
    nickname VARCHAR(25),
    birthday DATE NOT NULL
);
