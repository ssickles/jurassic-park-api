package mock

import "jurassic-park-api/data"

func NewStore() (data.Store, error) {
	return data.Store{
		Cages:     &Cages{},
		Dinosaurs: &Dinosaurs{},
	}, nil
}
