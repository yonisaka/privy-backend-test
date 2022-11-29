-- +goose Up
-- +goose StatementBegin
INSERT INTO cakes (`id`, `title`, `description`, `rating`, `image`, `created_at`, `updated_at`)
VALUES 
    (1, "Lemon cheesecake", "A cheesecake made of lemon", 7, "no-image.jpeg", NOW(), NOW()),
    (2, "Cake cake", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam", 5.5, "no-image.jpeg", NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM cakes;
-- +goose StatementEnd
