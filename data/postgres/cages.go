package postgres

import (
	"errors"
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

func (c Cages) Find(id int64) (*models.Cage, error) {
	var cage models.Cage
	err := c.Db.Model(&cage).
		Where("id = ?", id).
		Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &cage, nil
}

func (c Cages) FindByName(name string) (*models.Cage, error) {
	var cage models.Cage
	err := c.Db.Model(&cage).
		Where("name = ?", name).
		Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &cage, nil
}

func (c Cages) GetCageFoodType(id int64) (string, error) {
	var species models.Species
	err := c.Db.Model(&species).
		Join("JOIN dinosaurs ON species.name = dinosaurs.species_name").
		Where("dinosaurs.cage_id = ?", id).
		Column("food_type").
		First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return "", nil
		}
		return "", err
	}
	return species.FoodType, nil
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
