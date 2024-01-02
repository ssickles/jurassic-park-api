-- +goose Up
-- +goose StatementBegin
alter table dinosaurs add column cage_id bigint references cages(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table dinosaurs drop column cage_id;
-- +goose StatementEnd
