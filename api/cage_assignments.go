package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"jurassic-park-api/data"
	"jurassic-park-api/park"
	"log"
	"net/http"
	"strconv"
)

type CageAssignmentsController struct {
	Store data.Store
}

func (cac CageAssignmentsController) List(context *gin.Context) {
	cageIdInput := context.Param("cage_id")
	cageId, err := strconv.ParseInt(cageIdInput, 10, 64)
	if err != nil {
		log.Printf("error parsing cage id %s: %s\n", cageIdInput, err)
		context.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Invalid cage id provided"}})
		return
	}

	dinosaurs, err := cac.Store.Dinosaurs.FindByCageId(cageId)
	if err != nil {
		log.Printf("error listing the dinosaurs assigned to cage %d: %s\n", cageId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"errors": []string{"An unexpected error occurred"}})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": dinosaurs})
}

type CreateCageAssignmentPayload struct {
	DinosaurName string `json:"dinosaurName"`
}

func (cac CageAssignmentsController) Create(context *gin.Context) {
	cageIdInput := context.Param("cage_id")
	cageId, err := strconv.ParseInt(cageIdInput, 10, 64)
	if err != nil {
		log.Printf("error parsing cage id %s: %s\n", cageIdInput, err)
		context.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Invalid cage id provided"}})
		return
	}

	var payload CreateCageAssignmentPayload
	// First serialize the payload into a CreateCageAssignmentPayload struct
	err = context.BindJSON(&payload)
	if err != nil {
		log.Printf("error binding json for CreateCageAssignmentPayload: %s\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Invalid json provided for a cage assignment"}})
		return
	}

	err = park.CreateCageAssignment(cac.Store, cageId, payload.DinosaurName)
	if err != nil {
		switch {
		case errors.As(err, &park.CageNotFoundError{}):
			context.JSON(http.StatusBadRequest, gin.H{"errors": []string{err.Error()}})
			return
		case errors.As(err, &park.DinosaurNotFoundError{}):
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

	context.Status(http.StatusCreated)
}
