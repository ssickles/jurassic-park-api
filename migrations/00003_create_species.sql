-- +goose Up
-- +goose StatementBegin
create table if not exists species
(
    name        text        not null primary key,
    food_type   text        not null
);

alter table dinosaurs
    rename column species to species_name;

alter table dinosaurs
    add constraint fk_dinosaurs_species_name foreign key (species_name) references species (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists species;
-- +goose StatementEnd
