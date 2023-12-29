package postgres

import (
	"github.com/go-pg/pg/v10"
	"jurassic-park-api/data"
	"jurassic-park-api/models"
)

// Dinosaurs is a postgres implementation of data.Dinosaurs.
// This ensures that we follow the contract
var _ data.Dinosaurs = &Dinosaurs{}

type Dinosaurs struct {
	Db *pg.DB
}

func (d Dinosaurs) List() ([]models.Dinosaur, error) {
	var dinosaurs []models.Dinosaur
	err := d.Db.Model(&dinosaurs).Select()
	if err != nil {
		return nil, err
	}
	return dinosaurs, nil
}

func (d Dinosaurs) Create(dinosaur models.Dinosaur) (*models.Dinosaur, error) {
	var createdDinosaur models.Dinosaur
	_, err := d.Db.Model(&dinosaur).Returning("*").Insert(&createdDinosaur)
	if err != nil {
		return nil, err
	}
	return &createdDinosaur, nil
}
