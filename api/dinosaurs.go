package api

import (
	"github.com/gin-gonic/gin"
	"jurassic-park-api/data"
	"jurassic-park-api/models"
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
	var dinosaur models.Dinosaur
	err := context.BindJSON(&dinosaur)
	if err != nil {
		log.Printf("error binding dinosaur json: %s\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Invalid json provided for a dinosaur"}})
		return
	}

	species, err := dc.Store.Species.Find(dinosaur.SpeciesName)
	if err != nil {
		log.Printf("error getting species: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"errors": []string{"An unexpected error occurred"}})
		return
	}
	if species == nil {
		context.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Invalid species name provided"}})
		return
	}

	createdDinosaur, err := dc.Store.Dinosaurs.Create(dinosaur)
	if err != nil {
		log.Printf("error creating dinosaur: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"errors": []string{"An unexpected error occurred"}})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": createdDinosaur})
}
