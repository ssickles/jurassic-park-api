package data

import "jurassic-park-api/models"

type Store struct {
	Dinosaurs Dinosaurs
}

type Dinosaurs interface {
	List() ([]models.Dinosaur, error)
	Create(dinosaur models.Dinosaur) (*models.Dinosaur, error)
}
