-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists cakes (
	id int auto_increment NOT NULL,
	title varchar(255) NOT NULL,
	description TEXT NULL,
	rating FLOAT NULL,
	image varchar(255) NULL,
	created_at DATETIME NULL,
	updated_at DATETIME NULL,
	CONSTRAINT cakes_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cakes;
-- +goose StatementEnd
