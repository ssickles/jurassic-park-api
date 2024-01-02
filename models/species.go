package models

type Species struct {
	Name     string `json:"name" pg:"name,pk"`
	FoodType string `json:"foodType" pg:"food_type"`
}
