package mock

import "jurassic-park-api/data"

func NewStore() (data.Store, error) {
	species := &Species{}
	dinosaurs := &Dinosaurs{}

	return data.Store{
		Cages:     &Cages{Species: species, Dinosaurs: dinosaurs},
		Species:   species,
		Dinosaurs: dinosaurs,
	}, nil
}
