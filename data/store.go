package data

import "jurassic-park-api/models"

type Store struct {
	Cages     Cages
	Species   Species
	Dinosaurs Dinosaurs
}

type Cages interface {
	Find(id int64) (*models.Cage, error)
	FindByName(name string) (*models.Cage, error)
	GetCageFoodType(id int64) (string, error)
	List() ([]models.Cage, error)
	Create(cage models.Cage) (*models.Cage, error)
}

type Species interface {
	Find(name string) (*models.Species, error)
	Create(species models.Species) (*models.Species, error)
}

type Dinosaurs interface {
	FindByName(name string) (*models.Dinosaur, error)
	FindByCageId(cageId int64) ([]models.Dinosaur, error)
	List() ([]models.Dinosaur, error)
	Create(dinosaur models.Dinosaur) (*models.Dinosaur, error)
	Update(dinosaur models.Dinosaur) error
}
