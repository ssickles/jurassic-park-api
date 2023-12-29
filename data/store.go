package data

import "jurassic-park-api/models"

type Store struct {
	Cages     Cages
	Dinosaurs Dinosaurs
}

type Cages interface {
	List() ([]models.Cage, error)
	Create(cage models.Cage) (*models.Cage, error)
}

type Dinosaurs interface {
	List() ([]models.Dinosaur, error)
	Create(dinosaur models.Dinosaur) (*models.Dinosaur, error)
}
