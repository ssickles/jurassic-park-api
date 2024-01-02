package data

import "jurassic-park-api/models"

type Store struct {
	Cages     Cages
	Species   Species
	Dinosaurs Dinosaurs
}

type Cages interface {
	List() ([]models.Cage, error)
	Create(cage models.Cage) (*models.Cage, error)
}

type Species interface {
	Find(name string) (*models.Species, error)
	Create(species models.Species) (*models.Species, error)
}

type Dinosaurs interface {
	List() ([]models.Dinosaur, error)
	Create(dinosaur models.Dinosaur) (*models.Dinosaur, error)
}
