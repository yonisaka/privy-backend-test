-- +goose Up
-- +goose StatementBegin
INSERT INTO users (`id`, `username`, `email`, `password`, `created_at`, `updated_at`)
VALUES 
    (1, "userdemo", "user@gmail.com", "$2a$14$WjHP5kc8tE0LQzVuTqlwNehjNbYpUgz1f7hgv.98XWd8mgoJMMzFa", NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users;
-- +goose StatementEnd
