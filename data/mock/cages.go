package mock

import (
	"jurassic-park-api/data"
	"jurassic-park-api/models"
	"sync"
)

var _ data.Cages = &Cages{}

type Cages struct {
	primary map[int64]*models.Cage
	serial  int64
	init    sync.Once
}

func (c *Cages) ensure() {
	c.init.Do(func() {
		c.primary = map[int64]*models.Cage{}
	})
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
