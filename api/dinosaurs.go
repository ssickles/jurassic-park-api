package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"jurassic-park-api/data"
	"jurassic-park-api/park"
	"log"
	"net/http"
)

type DinosaursController struct {
	Store data.Store
}

func (dc DinosaursController) List(context *gin.Context) {
	dinosaurs, err := dc.Store.Dinosaurs.List()
	if err != nil {
		log.Printf("error listing dinosaurs: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"errors": []string{"An unexpected error occurred"}})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": dinosaurs})
}

func (dc DinosaursController) Create(context *gin.Context) {
	var payload park.CreateDinosaurPayload
	// First serialize the payload into a CreateDinosaurPayload struct
	err := context.BindJSON(&payload)
	if err != nil {
		log.Printf("error binding json for CreateDinosaurPayload: %s\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Invalid json provided for a dinosaur"}})
		return
	}

	// Then pass the payload to the CreateDinosaur function and allow our park package to handle the business logic
	createdDinosaur, err := park.CreateDinosaur(dc.Store, payload)
	if err != nil {
		// We will return a different error message depending on the type of error
		switch {
		case errors.As(err, &park.DinosaurNameAlreadyExistsError{}):
			context.JSON(http.StatusConflict, gin.H{"errors": []string{err.Error()}})
			return
		case errors.As(err, &park.InvalidSpeciesNameError{}):
			context.JSON(http.StatusBadRequest, gin.H{"errors": []string{err.Error()}})
			return
		case errors.As(err, &park.InvalidCageNameError{}):
			context.JSON(http.StatusBadRequest, gin.H{"errors": []string{err.Error()}})
			return
		case errors.As(err, &park.CageNotActiveError{}):
			context.JSON(http.StatusBadRequest, gin.H{"errors": []string{err.Error()}})
			return
		case errors.As(err, &park.CageAtCapacityError{}):
			context.JSON(http.StatusBadRequest, gin.H{"errors": []string{err.Error()}})
			return
		case errors.As(err, &park.MismatchedFoodTypeError{}):
			context.JSON(http.StatusBadRequest, gin.H{"errors": []string{err.Error()}})
			return
		default:
			log.Printf("%s\n", err)
			context.JSON(http.StatusInternalServerError, gin.H{"errors": []string{"An unexpected error occurred"}})
			return
		}
	}

	context.JSON(http.StatusCreated, gin.H{"data": createdDinosaur})
}
