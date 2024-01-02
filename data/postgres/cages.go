package postgres

import (
	"github.com/go-pg/pg/v10"
	"jurassic-park-api/data"
	"jurassic-park-api/models"
)

// Cages is a postgres implementation of data.Cages.
// This ensures that we follow the contract
var _ data.Cages = &Cages{}

type Cages struct {
	Db *pg.DB
}

func (c Cages) List() ([]models.Cage, error) {
	var cages []models.Cage
	err := c.Db.Model(&cages).
		Select()
	if err != nil {
		return nil, err
	}
	return cages, nil
}

func (c Cages) Create(cage models.Cage) (*models.Cage, error) {
	var createdCage models.Cage
	_, err := c.Db.Model(&cage).
		Returning("*").
		Insert(&createdCage)
	if err != nil {
		return nil, err
	}
	return &createdCage, nil
}
