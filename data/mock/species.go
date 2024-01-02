package mock

import (
	"jurassic-park-api/data"
	"jurassic-park-api/models"
	"sync"
)

var _ data.Species = &Species{}

type Species struct {
	primary map[string]*models.Species
	init    sync.Once
}

func (s *Species) ensure() {
	s.init.Do(func() {
		s.primary = map[string]*models.Species{}
	})
}

func (s *Species) Find(name string) (*models.Species, error) {
	s.ensure()
	for _, value := range s.primary {
		if value.Name == name {
			return value, nil
		}
	}
	return nil, nil
}

func (s *Species) Create(species models.Species) (*models.Species, error) {
	s.ensure()
	s.primary[species.Name] = &species
	return &species, nil
}
