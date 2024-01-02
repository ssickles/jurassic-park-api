package park

import (
	"fmt"
	"jurassic-park-api/data"
	"jurassic-park-api/models"
)

type CreateCagePayload struct {
	Name     string `json:"name" binding:"required"`
	Capacity int    `json:"capacity" binding:"required,number"`
}

type CageNameAlreadyExistsError struct {
	CageName string
}

func (e CageNameAlreadyExistsError) Error() string {
	return fmt.Sprintf("A cage with this name already exists: %s", e.CageName)
}

func CreateCage(store data.Store, payload CreateCagePayload) (*models.Cage, error) {
	cage := models.Cage{
		Name:     payload.Name,
		Capacity: payload.Capacity,
	}

	existing, err := store.Cages.FindByName(payload.Name)
	if err != nil {
		return nil, fmt.Errorf("error getting cage by name (%s): %w", payload.Name, err)
	}
	if existing != nil {
		return nil, CageNameAlreadyExistsError{CageName: payload.Name}
	}

	createdCage, err := store.Cages.Create(cage)
	if err != nil {
		return nil, fmt.Errorf("error creating cage: %w", err)
	}

	return createdCage, nil
}
