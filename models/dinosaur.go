package models

import "time"

type Dinosaur struct {
	Id        int64       `json:"id" pg:"id,pk"`
	Name      string      `json:"name" pg:"name" binding:"required"`
	Species   SpeciesName `json:"species" pg:"species" binding:"required"`
	CreatedAt time.Time   `json:"createdAt" pg:"created_at"`
}
