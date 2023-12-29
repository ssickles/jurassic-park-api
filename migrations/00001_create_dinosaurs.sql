-- +goose Up
-- +goose StatementBegin
create table if not exists dinosaurs
(
    id         bigserial   not null primary key,
    name       text        not null,
    species    text        not null,
    created_at timestamptz not null default now(),
    unique (name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists dinosaurs;
-- +goose StatementEnd
