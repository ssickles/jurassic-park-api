package mock

import (
	"jurassic-park-api/data"
	"jurassic-park-api/models"
	"sync"
)

var _ data.Cages = &Cages{}

type Cages struct {
	Species   data.Species
	Dinosaurs data.Dinosaurs
	primary   map[int64]*models.Cage
	serial    int64
	init      sync.Once
}

func (c *Cages) ensure() {
	c.init.Do(func() {
		c.primary = map[int64]*models.Cage{}
	})
}

func (c *Cages) Find(id int64) (*models.Cage, error) {
	c.ensure()
	return c.primary[id], nil
}

func (c *Cages) FindByName(name string) (*models.Cage, error) {
	c.ensure()
	for _, value := range c.primary {
		if value.Name == name {
			return value, nil
		}
	}
	return nil, nil
}

func (c *Cages) GetCageFoodType(id int64) (string, error) {
	c.ensure()
	dinosaurs, err := c.Dinosaurs.List()
	if err != nil {
		return "", err
	}
	for _, dinosaur := range dinosaurs {
		if dinosaur.CageId == id {
			species, err := c.Species.Find(dinosaur.SpeciesName)
			if err != nil {
				return "", err
			}
			return species.FoodType, nil
		}
	}
	return "", nil
}

func (c *Cages) List() ([]models.Cage, error) {
	c.ensure()
	var cages []models.Cage
	for _, value := range c.primary {
		cages = append(cages, *value)
	}
	return cages, nil
}

func (c *Cages) Create(cage models.Cage) (*models.Cage, error) {
	c.ensure()
	c.serial++
	cage.Id = c.serial
	c.primary[c.serial] = &cage
	return &cage, nil
}
