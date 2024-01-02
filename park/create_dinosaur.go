package park

import (
	"fmt"
	"jurassic-park-api/data"
	"jurassic-park-api/models"
)

type CreateDinosaurPayload struct {
	Name        string `json:"name" binding:"required"`
	SpeciesName string `json:"speciesName" binding:"required"`
	CageName    string `json:"cageName"`
}

type DinosaurNameAlreadyExistsError struct {
	DinosaurName string
}

func (e DinosaurNameAlreadyExistsError) Error() string {
	return fmt.Sprintf("A dinosaur with this name already exists: %s", e.DinosaurName)
}

type InvalidSpeciesNameError struct {
	SpeciesName string
}

func (e InvalidSpeciesNameError) Error() string {
	return fmt.Sprintf("Invalid species name: %s", e.SpeciesName)
}

type InvalidCageNameError struct {
	CageName string
}

func (e InvalidCageNameError) Error() string {
	return fmt.Sprintf("Invalid cage name: %s", e.CageName)
}

func CreateDinosaur(store data.Store, payload CreateDinosaurPayload) (*models.Dinosaur, error) {
	existing, err := store.Dinosaurs.FindByName(payload.Name)
	if err != nil {
		return nil, fmt.Errorf("error getting dinosaur by name (%s): %w", payload.Name, err)
	}
	if existing != nil {
		return nil, DinosaurNameAlreadyExistsError{DinosaurName: payload.Name}
	}

	species, err := store.Species.Find(payload.SpeciesName)
	if err != nil {
		return nil, fmt.Errorf("error getting species by name (%s): %w", payload.SpeciesName, err)
	}
	if species == nil {
		return nil, InvalidSpeciesNameError{SpeciesName: payload.SpeciesName}
	}

	dinosaur := models.Dinosaur{
		Name:        payload.Name,
		SpeciesName: payload.SpeciesName,
	}

	// A cage name is optional so we only need to do this if one was provided
	if len(payload.CageName) > 0 {
		// If a cage name was provided, we need to make sure that the cage exists
		cage, err := store.Cages.FindByName(payload.CageName)
		if err != nil {
			return nil, fmt.Errorf("error getting cage by name (%s): %w", payload.CageName, err)
		}
		if cage == nil {
			return nil, InvalidCageNameError{CageName: payload.CageName}
		}

		err = validateCageAssignment(store, *cage, *species)
		if err != nil {
			return nil, err
		}

		dinosaur.CageId = cage.Id
	}

	createdDinosaur, err := store.Dinosaurs.Create(dinosaur)
	if err != nil {
		return nil, fmt.Errorf("error creating dinosaur: %w", err)
	}

	return createdDinosaur, nil
}
