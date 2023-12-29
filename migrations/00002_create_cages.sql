-- +goose Up
-- +goose StatementBegin
create table if not exists cages
(
    id         bigserial   not null primary key,
    name       text        not null,
    created_at timestamptz not null default now(),
    unique (name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists cages;
-- +goose StatementEnd
