package park

import (
	"fmt"
	"jurassic-park-api/data"
)

type CageNotFoundError struct {
	CageId int64
}

func (e CageNotFoundError) Error() string {
	return fmt.Sprintf("Cage not found: %d", e.CageId)
}

type DinosaurNotFoundError struct {
	DinosaurName string
}

func (e DinosaurNotFoundError) Error() string {
	return fmt.Sprintf("Dinosaur not found: %s", e.DinosaurName)
}

func CreateCageAssignment(store data.Store, cageId int64, dinosaurName string) error {
	cage, err := store.Cages.Find(cageId)
	if err != nil {
		return fmt.Errorf("error getting cage by id (%d): %w", cageId, err)
	}
	if cage == nil {
		return CageNotFoundError{CageId: cageId}
	}

	dinosaur, err := store.Dinosaurs.FindByName(dinosaurName)
	if err != nil {
		return fmt.Errorf("error getting dinosaur by name (%s): %w", dinosaurName, err)
	}
	if dinosaur == nil {
		return DinosaurNotFoundError{DinosaurName: dinosaurName}
	}

	species, err := store.Species.Find(dinosaur.SpeciesName)
	if err != nil {
		return fmt.Errorf("error getting species by name (%s): %w", dinosaur.SpeciesName, err)
	}
	if species == nil {
		return fmt.Errorf("species not found: %s", dinosaur.SpeciesName)
	}

	err = validateCageAssignment(store, *cage, *species)
	if err != nil {
		return err
	}

	dinosaur.CageId = cageId
	err = store.Dinosaurs.Update(*dinosaur)
	if err != nil {
		return fmt.Errorf("error updating dinosaur: %w", err)
	}

	return nil
}
