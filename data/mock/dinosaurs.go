package mock

import (
	"jurassic-park-api/data"
	"jurassic-park-api/models"
	"sync"
)

var _ data.Dinosaurs = &Dinosaurs{}

type Dinosaurs struct {
	primary map[int64]*models.Dinosaur
	serial  int64
	init    sync.Once
}

func (d *Dinosaurs) ensure() {
	d.init.Do(func() {
		d.primary = map[int64]*models.Dinosaur{}
	})
}

func (d *Dinosaurs) FindByName(name string) (*models.Dinosaur, error) {
	d.ensure()
	for _, value := range d.primary {
		if value.Name == name {
			return value, nil
		}
	}
	return nil, nil
}

func (d *Dinosaurs) FindByCageId(cageId int64) ([]models.Dinosaur, error) {
	d.ensure()
	var dinosaurs []models.Dinosaur
	for _, value := range d.primary {
		if value.CageId == cageId {
			dinosaurs = append(dinosaurs, *value)
		}
	}
	return dinosaurs, nil
}

func (d *Dinosaurs) List() ([]models.Dinosaur, error) {
	d.ensure()
	var dinosaurs []models.Dinosaur
	for _, value := range d.primary {
		dinosaurs = append(dinosaurs, *value)
	}
	return dinosaurs, nil
}

func (d *Dinosaurs) Create(dinosaur models.Dinosaur) (*models.Dinosaur, error) {
	d.ensure()
	d.serial++
	dinosaur.Id = d.serial
	d.primary[d.serial] = &dinosaur
	return &dinosaur, nil
}

func (d *Dinosaurs) Update(dinosaur models.Dinosaur) error {
	d.ensure()
	d.primary[dinosaur.Id] = &dinosaur
	return nil
}
