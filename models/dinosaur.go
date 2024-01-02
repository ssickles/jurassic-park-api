package models

import "time"

type Dinosaur struct {
	Id          int64     `json:"id" pg:"id,pk"`
	Name        string    `json:"name" pg:"name"`
	SpeciesName string    `json:"speciesName" pg:"species_name"`
	CageId      int64     `json:"cageId" pg:"cage_id"`
	CreatedAt   time.Time `json:"createdAt" pg:"created_at"`
}
