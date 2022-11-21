-- +goose Up
CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    username      text UNIQUE NOT NULL,
    password_hash text        NOT NULL,
    created_at    timestamp DEFAULT NOW()
);

CREATE TABLE notes
(
    id          SERIAL PRIMARY KEY,
    uploader_id INT  NOT NULL,
    folder      text NOT NULL,
    title       text NOT NULL,
    data        text,
    public_id   text,
    updated_at  timestamp DEFAULT NOW()
);

-- +goose Down
DROP TABLE users;
DROP TABLE notes;
