-- +goose Up
-- +goose StatementBegin
alter table cages add column power_status text not null default 'ACTIVE';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table cages drop column power_status;
-- +goose StatementEnd
