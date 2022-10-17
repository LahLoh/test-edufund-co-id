-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id int UNSIGNED AUTO_INCREMENT,
    fullname varchar(255) NOT NULL,
    username varchar(255) NOT NULL,
    password text NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE IF EXISTS users;
