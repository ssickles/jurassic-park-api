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

type CageAtCapacityError struct {
	Capacity int
}

func (e CageAtCapacityError) Error() string {
	return fmt.Sprintf("The cage is already at capacity (%d), can't add another dinosaur", e.Capacity)
}

type MismatchedFoodTypeError struct {
	CageFoodType     string
	DinosaurFoodType string
}

func (e MismatchedFoodTypeError) Error() string {
	return fmt.Sprintf("The dinosaur's food type (%s) does not match the cage's food type (%s)", e.DinosaurFoodType, e.CageFoodType)
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

		inCage, err := store.Dinosaurs.FindByCageId(cage.Id)
		if err != nil {
			return nil, fmt.Errorf("error getting dinosaurs by cage id (%d): %w", cage.Id, err)
		}
		if len(inCage) >= cage.Capacity {
			return nil, CageAtCapacityError{Capacity: cage.Capacity}
		}

		// We also need to make sure the dinosaur's food type matches the cage's food type
		cageFoodType, err := store.Cages.GetCageFoodType(cage.Id)
		if err != nil {
			return nil, fmt.Errorf("error getting cage food type for cage %s: %w", cage.Name, err)
		}
		// If the cage doesn't have any dinosaurs, it won't have a food type yet and we can add any dinosaur
		if cageFoodType != "" && species.FoodType != cageFoodType {
			return nil, MismatchedFoodTypeError{CageFoodType: cageFoodType, DinosaurFoodType: species.FoodType}
		}

		dinosaur.CageId = cage.Id
	}

	createdDinosaur, err := store.Dinosaurs.Create(dinosaur)
	if err != nil {
		return nil, fmt.Errorf("error creating dinosaur: %w", err)
	}

	return createdDinosaur, nil
}
