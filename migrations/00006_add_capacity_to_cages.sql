-- +goose Up
-- +goose StatementBegin
alter table cages add column capacity int;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table cages drop column capacity;
-- +goose StatementEnd
