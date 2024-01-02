package park

import (
	"fmt"
	"jurassic-park-api/data"
	"jurassic-park-api/models"
)

func validateCageAssignment(store data.Store, cage models.Cage, species models.Species) error {
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
