package postgres

import (
	"errors"
	"github.com/go-pg/pg/v10"
	"jurassic-park-api/data"
	"jurassic-park-api/models"
)

var _ data.Species = &Species{}

type Species struct {
	Db *pg.DB
}

func (s Species) Find(name string) (*models.Species, error) {
	var species models.Species
	err := s.Db.Model(&species).
		Where(`name = ?`, name).
		First()
	if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}
	return &species, err
}

func (s Species) Create(species models.Species) (*models.Species, error) {
	_, err := s.Db.Model(&species).
		Insert()
	if err != nil {
		return nil, err
	}
	return &species, nil
}
