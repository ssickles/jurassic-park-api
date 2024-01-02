package park

import (
	"fmt"
	"jurassic-park-api/data"
	"jurassic-park-api/models"
)

type CreateCagePayload struct {
	Name        string             `json:"name" binding:"required"`
	Capacity    int                `json:"capacity" binding:"required,number"`
	PowerStatus models.PowerStatus `json:"powerStatus"`
}

type CageNameAlreadyExistsError struct {
	CageName string
}

func (e CageNameAlreadyExistsError) Error() string {
	return fmt.Sprintf("A cage with this name already exists: %s", e.CageName)
}

type InvalidCagePowerStatusError struct {
	PowerStatus models.PowerStatus
}

func (e InvalidCagePowerStatusError) Error() string {
	return fmt.Sprintf("Invalid cage power status: %s, must be ACTIVE or DOWN", e.PowerStatus)
}

func CreateCage(store data.Store, payload CreateCagePayload) (*models.Cage, error) {
	if payload.PowerStatus != "" && payload.PowerStatus != models.PowerStatusActive && payload.PowerStatus != models.PowerStatusDown {
		return nil, InvalidCagePowerStatusError{PowerStatus: payload.PowerStatus}
	}

	cage := models.Cage{
		Name:        payload.Name,
		Capacity:    payload.Capacity,
		PowerStatus: payload.PowerStatus,
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
