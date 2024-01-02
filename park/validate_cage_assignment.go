package park

import (
	"fmt"
	"jurassic-park-api/data"
	"jurassic-park-api/models"
)

type CageNotActiveError struct {
	CageName string
}

func (e CageNotActiveError) Error() string {
	return fmt.Sprintf("Can not assign a dinosaur to this cage (%s), it is not active", e.CageName)
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

func validateCageAssignment(store data.Store, cage models.Cage, species models.Species) error {
	if cage.PowerStatus != models.PowerStatusActive {
		return CageNotActiveError{CageName: cage.Name}
	}

	inCage, err := store.Dinosaurs.FindByCageId(cage.Id)
	if err != nil {
		return fmt.Errorf("error getting dinosaurs by cage id (%d): %w", cage.Id, err)
	}
	if len(inCage) >= cage.Capacity {
		return CageAtCapacityError{Capacity: cage.Capacity}
	}

	// We also need to make sure the dinosaur's food type matches the cage's food type
	cageFoodType, err := store.Cages.GetCageFoodType(cage.Id)
	if err != nil {
		return fmt.Errorf("error getting cage food type for cage %s: %w", cage.Name, err)
	}
	// If the cage doesn't have any dinosaurs, it won't have a food type yet and we can add any dinosaur
	if cageFoodType != "" && species.FoodType != cageFoodType {
		return MismatchedFoodTypeError{CageFoodType: cageFoodType, DinosaurFoodType: species.FoodType}
	}

	return nil
}
