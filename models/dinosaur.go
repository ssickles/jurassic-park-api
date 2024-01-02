package models

import "time"

type Dinosaur struct {
	Id          int64     `json:"id" pg:"id,pk"`
	Name        string    `json:"name" pg:"name,not_null" binding:"required"`
	SpeciesName string    `json:"speciesName" pg:"species_name,not_null" binding:"required"`
	CreatedAt   time.Time `json:"createdAt" pg:"created_at"`
}
