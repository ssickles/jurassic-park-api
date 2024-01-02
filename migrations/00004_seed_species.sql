-- +goose Up
-- +goose StatementBegin
/*
 * Pre-populate the species table with some data so we can get straight to creating dinosaurs.
 */
insert into species (name, food_type) values
    ('Tyrannosaurus', 'carnivore'),
    ('Velociraptor', 'carnivore'),
    ('Spinosaurus', 'carnivore'),
    ('Megalosaurus', 'carnivore'),
    ('Brachiosaurus', 'herbivore'),
    ('Stegosaurus', 'herbivore'),
    ('Ankylosaurus', 'herbivore'),
    ('Triceratops', 'herbivore');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
