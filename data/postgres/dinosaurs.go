package postgres

import (
	"errors"
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

func (d Dinosaurs) FindByName(name string) (*models.Dinosaur, error) {
	var dinosaur models.Dinosaur
	err := d.Db.Model(&dinosaur).
		Where("name = ?", name).
		Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &dinosaur, nil
}

func (d Dinosaurs) FindByCageId(cageId int64) ([]models.Dinosaur, error) {
	var dinosaurs []models.Dinosaur
	err := d.Db.Model(&dinosaurs).
		Where("cage_id = ?", cageId).
		Select()
	if err != nil {
		return nil, err
	}
	return dinosaurs, nil
}

func (d Dinosaurs) List() ([]models.Dinosaur, error) {
	var dinosaurs []models.Dinosaur
	err := d.Db.Model(&dinosaurs).
		Select()
	if err != nil {
		return nil, err
	}
	return dinosaurs, nil
}

func (d Dinosaurs) Create(dinosaur models.Dinosaur) (*models.Dinosaur, error) {
	var createdDinosaur models.Dinosaur
	_, err := d.Db.Model(&dinosaur).
		Returning("*").
		Insert(&createdDinosaur)
	if err != nil {
		return nil, err
	}
	return &createdDinosaur, nil
}
