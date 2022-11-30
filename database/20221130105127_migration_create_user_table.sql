-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists users (
	id int auto_increment NOT NULL,
	username varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	password varchar(255) NOT NULL,
	created_at DATETIME NULL,
	updated_at DATETIME NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
